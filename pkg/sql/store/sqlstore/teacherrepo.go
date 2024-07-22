package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeacherRepository struct {
	store *Store
}

func NewTeacherRepository(store *Store) *TeacherRepository {
	return &TeacherRepository{store: store}
}

func (r *TeacherRepository) LargerLinks(request map[int64][]int64) (map[*model.Group][]*model.Subject, error) {
	response := make(map[*model.Group][]*model.Subject)
	for groupID, subjectsID := range request {
		group, err := r.store.Group().Find(groupID)
		if err != nil {
			return nil, err
		}
		response[group] = make([]*model.Subject, 0)
		for _, subjectID := range subjectsID {
			subject, err := r.store.Subject().Find(subjectID)
			if err != nil {
				return nil, err
			}
			response[group] = append(response[group], subject)
		}
	}
	return response, nil
}
func (r *TeacherRepository) Create(teacher *model.Teacher) (*model.Teacher, error) {
	// Вставка данных учителя в MySQL

	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://admin:Erunda228@mongo")
	defer client.Disconnect(ctx)
	defer cancel()

	// Вставка данных Links в MongoDB
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	res, err := teacherLinksCollection.InsertOne(ctx, bson.M{"links": teacher.SL})
	if err != nil {
		return nil, err
	}

	teacher.LinksID = res.InsertedID.(primitive.ObjectID).Hex()

	query := "INSERT INTO teachers (name, capacity, links_id) VALUES (?, ?, ?)"
	result, err := r.store.db.Exec(query, teacher.Name, teacher.RecommendSchCap_, teacher.LinksID)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	teacher.ID = id

	return teacher, nil
}

func (r *TeacherRepository) Find(id int64) (*model.Teacher, error) {
	// Получение основных данных учителя из MySQL
	teacher := &model.Teacher{}
	err := r.store.db.QueryRow(
		"SELECT id, name, capacity, links_id FROM teachers WHERE id = ?",
		id,
	).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.RecommendSchCap_,
		&teacher.LinksID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, fmt.Errorf("err0 %s ", err.Error())
	}

	// Проверка корректности links_id
	linksID, err := primitive.ObjectIDFromHex(teacher.LinksID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID: %s", err.Error())
	}

	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()

	// Получение данных Links из MongoDB
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	var result bson.M
	err = teacherLinksCollection.FindOne(ctx, bson.M{"_id": linksID}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("err1 no documents in result: %s", err.Error())
		}
		return nil, fmt.Errorf("err1 %s ", err.Error())
	}

	// Преобразование данных
	links, err := utils.ConvertToSL(result)
	if err != nil {
		return nil, fmt.Errorf("failed to convert: %s", err.Error())
	}

	teacher.SL = links
	return teacher, nil
}

func (r *TeacherRepository) FindByName(name string) (*model.Teacher, error) {
	teacher := &model.Teacher{}
	err := r.store.db.QueryRow(
		"SELECT id, name, capacity, links_id FROM teachers WHERE name = ?",
		name,
	).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.RecommendSchCap_,
		&teacher.LinksID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return teacher, nil
}
func (r *TeacherRepository) GetList(page, limit int64) ([]*model.Teacher, error) {
	offset := (page - 1) * limit

	rows, err := r.store.db.Query(
		"SELECT id, name, capacity FROM teachers LIMIT ? OFFSET ?",
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*model.Teacher
	for rows.Next() {
		teacher := &model.Teacher{}
		if err := rows.Scan(
			&teacher.ID,
			&teacher.Name,
			&teacher.RecommendSchCap_,
		); err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (r *TeacherRepository) Update(teacher *model.Teacher) error {
	// Проверка наличия учителя
	old, err := r.store.teacherRepository.Find(teacher.ID)
	if err != nil {
		return fmt.Errorf("failed to find teacher: %s", err.Error())
	}

	// Проверка корректности links_id
	linksID, err := primitive.ObjectIDFromHex(old.LinksID)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %s", err.Error())
	}

	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()

	// Получение данных Links из MongoDB
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	var result bson.M
	err = teacherLinksCollection.FindOne(ctx, bson.M{"_id": linksID}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no documents found: %s", err.Error())
		}
		return fmt.Errorf("error finding documents: %s", err.Error())
	}

	// Преобразование данных
	links, err := utils.ConvertToSL(result)
	if err != nil {
		return fmt.Errorf("failed to convert: %s", err.Error())
	}

	// Сравнение и обновление данных, если они изменились
	if !utils.EqualMaps(links, teacher.SL) {
		// Обновление данных Links в MongoDB
		update := bson.M{"$set": bson.M{"links": teacher.SL}}
		_, err = teacherLinksCollection.UpdateOne(ctx, bson.M{"_id": linksID}, update)
		if err != nil {
			return fmt.Errorf("failed to update links: %s", err.Error())
		}
	}

	teacher.LinksID = old.LinksID

	// Обновление данных учителя в MySQL
	query, values := utils.UpdateTeachers(old, teacher)
	if len(values) == 0 {
		return nil
	}
	_, err = r.store.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
