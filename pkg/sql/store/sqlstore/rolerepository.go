package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

// RoleRepository struct
type RoleRepository struct {
	store *Store
}

var (
// errNotActive error = errors.New("user is not activated")
// errIncorrectParam error = errors.New("incorrect parameters to use")
)

// Initialization

func (rr *RoleRepository) CreateRole(name string) (*model.Role, error) {
	r := &model.Role{}
	result, err := rr.store.db.Exec("insert into roles (name) values (?)", name)
	if err != nil {
		return nil, fmt.Errorf("database role error:%s", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("database role error:%s", err.Error())
	}
	r.ID = int32(id)
	r.Name = name

	return r, nil

}

func (rr *RoleRepository) Find(id int64) (*model.Role, error) {
	r := &model.Role{}
	err := rr.store.db.QueryRow(
		"SELECT id, name FROM roles WHERE id = ?",
		id,
	).Scan(
		&r.ID,
		&r.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database role error:%s", store.ErrRec404.Error())
		}
		return nil, fmt.Errorf("database role error:%s", err.Error())

	}

	return r, nil
}

func (rr *RoleRepository) FindByName(name string) (*model.Role, error) {

	r := &model.Role{}

	err := rr.store.db.QueryRow(
		"SELECT id, name FROM roles WHERE name = ?", name).Scan(
		&r.ID,
		&r.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("database role error:%s", store.ErrRec404.Error())
		}
		return nil, fmt.Errorf("database role error:%s", err.Error())

	}
	return r, nil
}

func (r *RoleRepository) GetList(page int64, limit int64) (roles []*model.Role, err error) {
	rows, err := r.store.db.Query(
		"SELECT id, name FROM roles LIMIT ? OFFSET ?",
		limit,
		(page-1)*limit,
	)
	if err != nil {
		return nil, fmt.Errorf("database role error:%s", err.Error())
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		role := &model.Role{}
		if err := rows.Scan(
			&role.ID,
			&role.Name,
		); err != nil {
			return nil, fmt.Errorf("database role error:%s", err.Error())
		}
		roles = append(roles, role)
	}

	return roles, nil
}
