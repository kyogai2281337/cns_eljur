package sqlstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

type SubjectRepository struct {
	store *Store
}

func NewSubjectRepository(store *Store) *SubjectRepository {
	return &SubjectRepository{store: store}
}

func (r *SubjectRepository) Create(ctx context.Context, subject *model.Subject) (*model.Subject, error) {
	tx, err := r.store.getTxFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO subjects (name, type) VALUES (?, ?)"
	result, err := tx.Exec(query, subject.Name, subject.RecommendCabType)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	subject.ID = id
	return subject, nil
}

func (r *SubjectRepository) Find(id int64) (*model.Subject, error) {
	subject := &model.Subject{}
	err := r.store.db.QueryRow(
		"SELECT id, name, type FROM subjects WHERE id = ?",
		id,
	).Scan(
		&subject.ID,
		&subject.Name,
		&subject.RecommendCabType,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return subject, nil
}

func (r *SubjectRepository) FindByName(name string) (*model.Subject, error) {
	subject := &model.Subject{}
	err := r.store.db.QueryRow(
		"SELECT id, name, type FROM subjects WHERE name = ?",
		name,
	).Scan(
		&subject.ID,
		&subject.Name,
		&subject.RecommendCabType,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return subject, nil
}

func (r *SubjectRepository) GetList(page int64, limit int64) ([]*model.Subject, error) {
	offset := (page - 1) * limit // Calculate offset for pagination
	rows, err := r.store.db.Query(
		"SELECT id, name, type FROM subjects LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subjects := make([]*model.Subject, 0)
	for rows.Next() {
		subject := &model.Subject{}
		if err := rows.Scan(&subject.ID, &subject.Name, &subject.RecommendCabType); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

func (r *SubjectRepository) Update(ctx context.Context, subject *model.Subject) error {
	_, err := r.Find(subject.ID)
	if err != nil {
		return err
	}

	tx, err := r.store.getTxFromCtx(ctx)
	if err != nil {
		return err
	}

	query := "UPDATE subjects SET name = ?, type = ? WHERE id = ?"
	_, err = tx.Exec(query, subject.Name, subject.RecommendCabType, subject.ID)
	if err != nil {
		return err
	}
	return nil
}
