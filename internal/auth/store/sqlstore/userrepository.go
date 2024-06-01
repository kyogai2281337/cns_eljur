package sqlstore

import (
	"database/sql"
	"errors"
	"github.com/kyogai2281337/cns_eljur/internal/auth/model"
	"github.com/kyogai2281337/cns_eljur/internal/auth/store"
	"log"
)

// UserRepository UserRep struct
type UserRepository struct {
	store *Store
}

var (
	errNotActive      error = errors.New("user is not activated")
	errIncorrectParam error = errors.New("incorrect parameters to use")
)

// Create Initialization
func (r *UserRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	result, err := r.store.db.Exec(
		"INSERT INTO users (email, encrypted_password, first_name, last_name, role_id) VALUES (?, ?, ?, ?, ?)",
		u.Email, u.EncPass, u.FirstName, u.LastName, u.Role.ID,
	)
	if err != nil {
		return err
	}

	// Получение ID вставленной записи
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	ok, err := r.CheckActive(email)
	if !ok || err != nil {
		return nil, errNotActive
	}
	var roleId int64
	err = r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, first_name, last_name, role_id FROM users WHERE email = ?",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncPass,
		&u.FirstName,
		&u.LastName,
		&roleId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	u.Role, err = r.store.Role().FindRoleById(roleId)
	if err != nil {
		return nil, err
	}
	err, u.PermsSet = r.SearchPermissions(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) CheckActive(param interface{}) (bool, error) {
	u := &model.User{}
	var row string
	switch param.(type) {
	case int64:
		row = "SELECT is_active FROM users WHERE id = ?"
	case string:
		row = "SELECT is_active FROM users WHERE email = ?"
	default:
		return false, errIncorrectParam
	}

	err := r.store.db.QueryRow(row, param).Scan(&u.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, store.ErrRec404
		}
		return false, err
	}

	return u.IsActive, nil
}

func (r *UserRepository) Find(id int64) (*model.User, error) {
	u := &model.User{}
	ok, err := r.CheckActive(id)
	if !ok || err != nil {
		return nil, errNotActive
	}
	var roleId int64
	err = r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, first_name, last_name, role_id FROM users WHERE id = ?",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncPass,
		&u.FirstName,
		&u.LastName,
		&roleId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRec404
		}
		return nil, err
	}
	u.Role, err = r.store.Role().FindRoleById(roleId)
	if err != nil {
		return nil, err
	}
	err, u.PermsSet = r.SearchPermissions(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Delete(id int64) error {
	_, err := r.store.db.Exec("update users set is_active = 0 where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (pp *UserRepository) SearchPermissions(u *model.User) (error, *[]model.Permission) {
	var permset []model.Permission
	query := "SELECT id_perm FROM usr_perms WHERE id_user = ?"

	rows, err := pp.store.db.Query(query, u.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		r := model.Permission{}
		var idPermission int32
		if err := rows.Scan(&idPermission); err != nil {
			return err, nil
		}
		err = pp.store.db.QueryRow(
			"SELECT id, name FROM permission WHERE id = ?",
			idPermission,
		).Scan(
			&r.Id,
			&r.Name,
		)
		permset = append(permset, r)
	}

	u.PermsSet = &permset
	return nil, u.PermsSet
}
