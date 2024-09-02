package sqlstore

import (
	"context"
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

func (g *GroupRepository) Create(ctx context.Context, query *model.Group) (*model.Group, error) {

	tx, err := g.store.getTxFromCtx(ctx)
	if err != nil {
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}

	result, err := tx.Exec("insert into `groups` (name, spec_id, max_pairs) values (?, ?, ?)", query.Name, query.Specialization.ID, query.MaxPairs)
	if err != nil {
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}
	query.ID = id
	query.Specialization, err = g.store.Specialization().Find(query.Specialization.ID)
	if err != nil {
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}
	return query, nil
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
	//fmt.Println("1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database group error:%s", store.ErrRec404.Error())
		}
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}
	//fmt.Println("2")
	group.Specialization, err = g.store.Specialization().Find(specId)
	if err != nil {
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}
	//fmt.Println("3")
	return group, nil
}

func (g *GroupRepository) FindByName(name string) (*model.Group, error) {
	group := &model.Group{}
	var specId int64
	err := g.store.db.QueryRow(
		"SELECT id, name, spec_id, max_pairs FROM `groups` WHERE name = ?",
		name,
	).Scan(
		&group.ID,
		&group.Name,
		&specId,
		&group.MaxPairs,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database group error:%s", store.ErrRec404.Error())
		}
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}

	group.Specialization, err = g.store.Specialization().Find(specId)
	if err != nil {
		return nil, fmt.Errorf("database group error:%s", err.Error())
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
		return nil, fmt.Errorf("database group error:%s", err.Error())
	}
	defer rows.Close()
	groups := make([]*model.Group, 0)
	for rows.Next() {
		group := &model.Group{}
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, fmt.Errorf("database group error:%s", err.Error())
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (g *GroupRepository) Update(ctx context.Context, group *model.Group) error {
	current, err := g.Find(group.ID)
	if err != nil {
		return fmt.Errorf("database group error:%s", err.Error())
	}
	query, values := utils.UpdateGroups(current, group)
	if len(values) == 0 {
		return nil
	}

	tx, err := g.store.getTxFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("database group error:%s", err.Error())
	}

	_, err = tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("database group error:%s", err.Error())
	}
	return nil
}
