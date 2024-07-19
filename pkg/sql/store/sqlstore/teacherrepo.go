package sqlstore

import (
	"database/sql"
	"errors"
	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()

	// Вставка данных Links в MongoDB
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	_, err = teacherLinksCollection.InsertOne(ctx, bson.M{"_id": teacher.LinksID, "links": teacher.ShorterLinks()})
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer client.Disconnect(ctx)
	defer cancel()

	// Получение данных Links из MongoDB
	teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
	sLinks := make(map[int64][]int64)
	err = teacherLinksCollection.FindOne(ctx, bson.M{"_id": teacher.LinksID}).Decode(&sLinks)
	if err != nil {
		return nil, err
	}

	// Преобразование данных Links
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
	offset := (page - 1) * limit
	rows, err := r.store.db.Query(
		"SELECT id, name, capacity, links_id FROM teachers LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	teachers := make([]*model.Teacher, 0)
	for rows.Next() {
		teacher := &model.Teacher{}
		if err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.RecommendSchCap_, &teacher.LinksID); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
		defer client.Disconnect(ctx)
		defer cancel()

		teacherLinksCollection := client.Database("eljur").Collection("teacher_links")
		sLinks := make(map[int64][]int64)
		err = teacherLinksCollection.FindOne(ctx, bson.M{"_id": teacher.LinksID}).Decode(&sLinks)
		if err != nil {
			log.Printf("Error querying MongoDB: %v", err)
			return nil, err
		}

		teacher.Links, err = r.LargerLinks(sLinks)
		if err != nil {
			log.Printf("Error transforming links: %v", err)
			return nil, err
		}

		teachers = append(teachers, teacher)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Rows error: %v", err)
		return nil, err
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
