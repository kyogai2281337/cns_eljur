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

type CabinetRepository struct {
	store *Store
}

func (c *CabinetRepository) Create(ctx context.Context, cabinet *model.Cabinet) (*model.Cabinet, error) {
	tx, err := c.store.getTxFromCtx(ctx)
	if err != nil {
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}

	result, err := tx.Exec("insert into cabinets (name, type, capacity) values (?, ?, ?)", cabinet.Name, cabinet.Type, cabinet.Capacity)
	if err != nil {
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}
	cabinet.ID = id
	return cabinet, nil
}

func (c *CabinetRepository) Find(id int64) (*model.Cabinet, error) {
	cab := &model.Cabinet{}
	err := c.store.db.QueryRow(
		"SELECT id, name, type, capacity FROM cabinets WHERE id = ?",
		id,
	).Scan(
		&cab.ID,
		&cab.Name,
		&cab.Type,
		&cab.Capacity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database cabinet error:%s", store.ErrRec404.Error())
		}
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}
	return cab, nil
}

func (c *CabinetRepository) FindByName(name string) (*model.Cabinet, error) {
	cab := &model.Cabinet{}
	err := c.store.db.QueryRow(
		"SELECT id, name, type, capacity FROM cabinets WHERE name = ?",
		name,
	).Scan(
		&cab.ID,
		&cab.Name,
		&cab.Type,
		&cab.Capacity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database cabinet error:%s", store.ErrRec404.Error())
		}
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}
	return cab, nil
}

func (c *CabinetRepository) GetList(page int64, limit int64) ([]*model.Cabinet, error) {
	offset := (page - 1) * limit // Calculate offset for pagination

	rows, err := c.store.db.Query(
		"SELECT id, name FROM cabinets LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr // Assign the error to the named return value
		}
	}()

	var cabs []*model.Cabinet
	for rows.Next() {
		u := &model.Cabinet{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
		); err != nil {
			return nil, fmt.Errorf("database cabinet error:%s", err.Error())
		}
		cabs = append(cabs, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("database cabinet error:%s", err.Error())
	}

	return cabs, nil
}

func (c *CabinetRepository) Update(ctx context.Context, cabinet *model.Cabinet) error {
	current, err := c.Find(cabinet.ID)
	if err != nil {
		return fmt.Errorf("database cabinet error:%s", err.Error())
	}
	query, values := utils.UpdateCabs(current, cabinet)
	if len(values) == 0 {
		return nil
	}

	tx, err := c.store.getTxFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("database cabinet error:%s", err.Error())
	}

	_, err = tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("database cabinet error:%s", err.Error())
	}
	return nil
}
