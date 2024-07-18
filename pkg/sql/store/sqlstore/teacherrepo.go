package sqlstore

import (
	"database/sql"
	"errors"

	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"go.mongodb.org/mongo-driver/bson"
)

type TeacherRepository struct {
	store *Store
}

func NewTeacherRepository(store *Store) *TeacherRepository {
	return &TeacherRepository{store: store}
}

func (t *TeacherRepository) LargerLinks(request map[int64][]int64) (map[*model.Group][]*model.Subject, error) {
	response := make(map[*model.Group][]*model.Subject)
	for groupID, subjectsID := range request {
		group, err := t.store.Group().Find(groupID)
		if err != nil {
			return nil, err
		}
		response[group] = make([]*model.Subject, 0)
		for _, subjectID := range subjectsID {
			subject, err := t.store.Subject().Find(subjectID)
			if err != nil {
				return nil, err
			}
			response[group] = append(response[group], subject)
		}

	}
	return response, nil
}
func (r *TeacherRepository) Create(teacher *model.Teacher) (*model.Teacher, error) {
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
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	_, err = teacherLinksCollection.InsertOne(ctx, teacher.ShorterLinks())
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (r *TeacherRepository) Find(id int64) (*model.Teacher, error) {
	teacher := &model.Teacher{}
	err := r.store.db.QueryRow(
		"SELECT id, name, capacity, links_id FROM `teachers` WHERE id = ?",
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
		return nil, err
	}
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	sLinks := make(map[int64][]int64)
	teacherLinksCollection.FindOne(ctx, bson.M{"_id": teacher.LinksID}).Decode(&sLinks)
	teacher.Links, err = r.LargerLinks(sLinks)
	if err != nil {
		return nil, err
	}
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

func (r *TeacherRepository) GetList(page int64, limit int64) ([]*model.Teacher, error) {
	offset := (page - 1) * limit // Calculate offset for pagination
	rows, err := r.store.db.Query(
		"SELECT id, name, capacity FROM teachers LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	teachers := make([]*model.Teacher, 0)
	for rows.Next() {
		teacher := &model.Teacher{}
		if err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.RecommendSchCap_); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func (r *TeacherRepository) Update(teacher *model.Teacher) error {
	_, err := r.Find(teacher.ID)
	if err != nil {
		return err
	}
	query := "UPDATE teachers SET name = ?, capacity = ?, links_id = ? WHERE id = ?"
	_, err = r.store.db.Exec(query, teacher.Name, teacher.RecommendSchCap_, teacher.LinksID, teacher.ID)
	if err != nil {
		return err
	}
	return nil
}
