package sqlstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpecializationRepository struct {
	store *Store
}

func (s *SpecializationRepository) LargerPlans(request map[int64]int) (map[*model.Subject]int, error) {
	response := make(map[*model.Subject]int)
	for subjectID, lessonsCount := range request {
		subject, err := s.store.Subject().Find(subjectID)
		if err != nil {
			return nil, err
		}
		response[subject] = lessonsCount
	}
	return response, nil
}

func (s *SpecializationRepository) Find(id int64) (*model.Specialization, error) {
	spec := &model.Specialization{}
	err := s.store.db.QueryRow("SELECT * FROM specializations WHERE id = ?", id).Scan(&spec.ID, &spec.Name, &spec.Course, &spec.PlanId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}

	planId, _ := primitive.ObjectIDFromHex(spec.PlanId)

	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://admin:Erunda228@mongo")
	defer client.Disconnect(ctx)
	defer cancel()

	SpecPlansCollection := client.Database("eljur").Collection("specialization_plans")
	var result bson.M
	err = SpecPlansCollection.FindOne(ctx, bson.M{"_id": planId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	spec.ShortPlan, err = utils.ConvertToPlan(result)

	if err != nil {
		return nil, err
	}

	spec.EduPlan, err = s.LargerPlans(spec.ShortPlan)
	if err != nil {
		return nil, fmt.Errorf("failed to convert plan: %s", err.Error())
	}

	return spec, nil
}

func (s *SpecializationRepository) Create(ctx context.Context, spec *model.Specialization) (*model.Specialization, error) {
	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://admin:Erunda228@mongo")
	defer client.Disconnect(ctx)
	defer cancel()

	// Вставка данных Links в MongoDB
	SpecPlansCollection := client.Database("eljur").Collection("specialization_plans")
	res, err := SpecPlansCollection.InsertOne(ctx, bson.M{"plans": spec.ShortPlan})
	if err != nil {
		return nil, err
	}

	spec.PlanId = res.InsertedID.(primitive.ObjectID).Hex()

	tx, err := s.store.getTxFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO specializations (name, course, plan_id) VALUES (?, ?, ?)", spec.Name, spec.Course, spec.PlanId)
	if err != nil {
		return nil, err
	}

	spec.EduPlan, err = s.LargerPlans(spec.ShortPlan)
	if err != nil {
		return nil, fmt.Errorf("failed to convert plan: %s", err.Error())
	}

	return spec, nil
}

func (s *SpecializationRepository) GetList(page int64, limit int64) ([]*model.Specialization, error) {
	offset := (page - 1) * limit
	rows, err := s.store.db.Query(
		"SELECT id, name FROM specializations LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	groups := make([]*model.Specialization, 0)
	for rows.Next() {
		group := &model.Specialization{}
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (s *SpecializationRepository) FindByName(name string) (*model.Specialization, error) {
	spec := &model.Specialization{}
	err := s.store.db.QueryRow("SELECT * FROM specializations WHERE name = ?", name).Scan(&spec.ID, &spec.Name, &spec.Course, &spec.PlanId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}

	planId, _ := primitive.ObjectIDFromHex(spec.PlanId)

	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://admin:Erunda228@mongo")
	defer client.Disconnect(ctx)
	defer cancel()

	SpecPlansCollection := client.Database("eljur").Collection("specialization_plans")
	var result bson.M
	err = SpecPlansCollection.FindOne(ctx, bson.M{"_id": planId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	spec.ShortPlan, err = utils.ConvertToPlan(result)

	if err != nil {
		return nil, err
	}

	spec.EduPlan, err = s.LargerPlans(spec.ShortPlan)
	if err != nil {
		return nil, fmt.Errorf("failed to convert plan: %s", err.Error())
	}

	return spec, nil
}

func (s *SpecializationRepository) Update(ctx context.Context, spec *model.Specialization) error {
	tx, err := s.store.getTxFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %s", err.Error())
	}
	// Проверка наличия spec
	old, err := s.store.specializationRepository.Find(spec.ID)
	if err != nil {
		return fmt.Errorf("failed to find specialization: %s", err.Error())
	}

	// Проверка корректности links_id
	planId, err := primitive.ObjectIDFromHex(old.PlanId)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %s", err.Error())
	}
	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()

	// Получение данных ShortPlan из MongoDB
	specShortPlanCollection := client.Database("eljur").Collection("specialization_plans")
	var result bson.M
	err = specShortPlanCollection.FindOne(ctx, bson.M{"_id": planId}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no documents found: %s", err.Error())
		}
		return fmt.Errorf("error finding documents: %s", err.Error())
	}

	// Преобразование данных
	plans, err := utils.ConvertToPlan(result)
	if err != nil {
		return fmt.Errorf("failed to convert: %s", err.Error())
	}

	if !utils.EqualEasyMaps(plans, spec.ShortPlan) {
		// Обновление данных Links в MongoDB
		update := bson.M{"$set": bson.M{"plans": plans, "_id": planId}}
		_, err = specShortPlanCollection.UpdateOne(ctx, bson.M{"_id": planId}, update)
		if err != nil {
			return fmt.Errorf("failed to update links: %s", err.Error())
		}
	}

	_, err = tx.Exec("UPDATE specializations SET name = ?, course = ? WHERE id = ?", spec.Name, spec.Course, spec.ID)
	if err != nil {
		return err
	}

	spec.EduPlan, err = s.LargerPlans(spec.ShortPlan)
	if err != nil {
		return fmt.Errorf("failed to convert plan: %s", err.Error())
	}

	spec.PlanId = old.PlanId
	return nil
}
