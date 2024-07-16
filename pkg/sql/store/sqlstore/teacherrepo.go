package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

type TeacherRepository struct {
	store *Store
}

func NewTeacherRepository(store *Store) *TeacherRepository {
	return &TeacherRepository{store: store}
}

func (r *TeacherRepository) Create(teacher *model.Teacher) (*model.Teacher, error) {
	query := "INSERT INTO teachers (name, capacity) VALUES (?, ?)"
	result, err := r.store.db.Exec(query, teacher.Name, teacher.RecommendSchCap_)
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
	teacher := &model.Teacher{}
	err := r.store.db.QueryRow(
		"SELECT id, name, capacity FROM teachers WHERE id = ?",
		id,
	).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.RecommendSchCap_,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return teacher, nil
}

func (r *TeacherRepository) FindByName(name string) (*model.Teacher, error) {
	teacher := &model.Teacher{}
	err := r.store.db.QueryRow(
		"SELECT id, name, capacity FROM teachers WHERE name = ?",
		name,
	).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.RecommendSchCap_,
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
	query := "UPDATE teachers SET name = ?, capacity = ? WHERE id = ?"
	_, err = r.store.db.Exec(query, teacher.Name, teacher.RecommendSchCap_, teacher.ID)
	if err != nil {
		return err
	}
	return nil
}
