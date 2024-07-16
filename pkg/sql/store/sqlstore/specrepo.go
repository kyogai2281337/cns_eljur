package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

type SpecializationRepository struct {
	store *Store
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
	return spec, nil
}

func (s *SpecializationRepository) Create(spec *model.Specialization) (*model.Specialization, error) {
	_, err := s.store.db.Exec("INSERT INTO specializations (name, course, plan_id) VALUES (?, ?, ?)", spec.Name, spec.Course, spec.PlanId)
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func (s *SpecializationRepository) GetList(page int64, limit int64) ([]*model.Specialization, error) {
	offset := (page - 1) * limit // Calculate offset for pagination
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
	return spec, nil
}

func (s *SpecializationRepository) Update(spec *model.Specialization) error {
	_, err := s.store.db.Exec("UPDATE specializations SET name = ?, course = ?, plan_id = ? WHERE id = ?", spec.Name, spec.Course, spec.PlanId, spec.ID)
	if err != nil {
		return err
	}
	return nil
}
