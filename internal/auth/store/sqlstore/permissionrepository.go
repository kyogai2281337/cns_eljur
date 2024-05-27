package sqlstore

import (
	"database/sql"
	"github.com/kyogai2281337/cns_eljur/internal/auth/model"
	"github.com/kyogai2281337/cns_eljur/internal/auth/store"
)

type PermissionRepository struct {
	store *Store
}

func (pp *PermissionRepository) IsAdmin(id int64) (bool, error) {
	p := &model.Permission{}
	err := pp.store.db.QueryRow(
		"SELECT id, isAdmin FROM roles WHERE id = ?",
		id,
	).Scan(
		&p.Id,
		&p.IsAdmin,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, store.ErrRec404
		}
		return false, err
	}

	return p.IsAdmin, nil
}
