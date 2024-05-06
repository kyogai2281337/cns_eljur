package sqlstore

import (
	"database/sql"

	"github.com/kyogai2281337/cns_eljur/internal/auth/model"
	"github.com/kyogai2281337/cns_eljur/internal/auth/store"
)

// UserRep struct
type RoleRepository struct {
	store *Store
}

var (
// errNotActive error = errors.New("user is not activated")
// errIncorrectParam error = errors.New("incorrect parameters to use")
)

// Initialization

func (rr *RoleRepository) Create(name string) (*model.Role, error) {
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
		if err == sql.ErrNoRows {
			return nil, store.ErrRec404
		}
		return nil, err
	}

	return r, nil
}