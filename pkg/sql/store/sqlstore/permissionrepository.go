package sqlstore

import (
	"database/sql"
	"errors"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

type PermissionRepository struct {
	store *Store
}

func (p *PermissionRepository) CreatePermission(name string, endpoint string) error {
	return nil
}

func (r *PermissionRepository) FindPermById(id int32) (*model.Permission, error) {
	p := &model.Permission{}
	err := r.store.db.QueryRow(
		"SELECT id, name FROM permission WHERE id = ?",
		id,
	).Scan(
		&p.Id,
		&p.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	return p, nil
}
