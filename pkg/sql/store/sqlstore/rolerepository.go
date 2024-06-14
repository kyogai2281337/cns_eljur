package sqlstore

import (
	"database/sql"
	"errors"

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
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	r.ID = int32(id)

	return r, nil

}

func (rr *RoleRepository) FindRoleById(id int64) (*model.Role, error) {
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
			return nil, store.ErrRec404
		}
		return nil, err
	}

	return r, nil
}

func (rr *RoleRepository) FindRoleByName(name string) (*model.Role, error) {

	r := &model.Role{}

	err := rr.store.db.QueryRow(
		"SELECT id, name FROM roles WHERE name = ?", name).Scan(
		&r.ID,
		&r.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return r, nil

}
