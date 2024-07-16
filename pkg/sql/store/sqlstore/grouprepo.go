package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore/utils"
)

type GroupRepository struct {
	store *Store
}

func (g *GroupRepository) Create(query *model.Group) (*model.Group, error) {
	group := &model.Group{}
	result, err := g.store.db.Exec("insert into `groups` (name, spec_id, max_pairs) values (?, ?, ?)", query.Name, query.Specialization.ID, query.MaxPairs)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	group.ID = id
	group.Specialization, err = g.store.Specialization().Find(query.Specialization.ID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GroupRepository) Find(id int64) (*model.Group, error) {
	group := &model.Group{}
	var specId int64
	err := g.store.db.QueryRow(
		"SELECT id, name, spec_id, max_pairs FROM `groups` WHERE id = ?",
		id,
	).Scan(
		&group.ID,
		&group.Name,
		&specId,
		&group.MaxPairs,
	)
	fmt.Println("1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	fmt.Println("2")
	group.Specialization, err = g.store.Specialization().Find(specId)
	if err != nil {
		return nil, err
	}
	fmt.Println("3")
	return group, nil
}

func (g *GroupRepository) FindByName(name string) (*model.Group, error) {
	group := &model.Group{}
	err := g.store.db.QueryRow(
		"SELECT id, name, spec_id, max_pairs FROM `groups` WHERE name = ?",
		name,
	).Scan(
		&group.ID,
		&group.Name,
		&group.Specialization.ID,
		&group.MaxPairs,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return group, nil
}

func (g *GroupRepository) GetList(page int64, limit int64) ([]*model.Group, error) {
	offset := (page - 1) * limit // Calculate offset for pagination
	rows, err := g.store.db.Query(
		"SELECT id, name FROM `groups` LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	groups := make([]*model.Group, 0)
	for rows.Next() {
		group := &model.Group{}
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (g *GroupRepository) Update(group *model.Group) error {
	current, err := g.Find(group.ID)
	if err != nil {
		return err
	}
	query, values := utils.UpdateGroups(current, group)
	if len(values) == 0 {
		return nil
	}
	_, err = g.store.db.Exec(query, values...)
	if err != nil {
		return err
	}
	return err
}
