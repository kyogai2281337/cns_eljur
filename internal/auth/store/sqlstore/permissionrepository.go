package sqlstore

import (
	"github.com/kyogai2281337/cns_eljur/internal/auth/model"
	"log"
)

type PermissionRepository struct {
	store *Store
}

func (p *PermissionRepository) CreatePermission(name string, endpoint string) error {
	return nil
}

func (pp *PermissionRepository) SearchPermissions(u *model.User) error {
	var permset []model.Permission
	query := "SELECT id, name FROM user_perms WHERE id = ?"

	rows, err := pp.store.db.Query(query, u.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		r := model.Permission{}
		var idPermission int
		if err := rows.Scan(&idPermission); err != nil {
			log.Fatal(err)
		}
		err = pp.store.db.QueryRow(
			"SELECT id, name FROM permissions WHERE id = ?",
			idPermission,
		).Scan(
			&r.Id,
			&r.Name,
		)
		permset = append(permset, r)
	}
	u.PermsSet = &permset
	return nil
}
