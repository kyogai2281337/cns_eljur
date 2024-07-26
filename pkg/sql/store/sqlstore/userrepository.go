package sqlstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore/utils"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

// UserRepository представляет репозиторий пользователей.
type UserRepository struct {
	store *Store
}

// Определение ошибок.
var (
	errNotActive      = errors.New("user is not activated")
	errIncorrectParam = errors.New("incorrect parameters to use")
)

// Create создает нового пользователя.
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	tx, err := r.store.getTxFromCtx(ctx)
	if err != nil {
		return err
	}

	result, err := tx.Exec(
		"INSERT INTO users (email, encrypted_password, first_name, last_name, role_id) VALUES (?, ?, ?, ?, ?)",
		u.Email, u.EncPass, u.FirstName, u.LastName, u.Role.ID,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id

	return nil
}

// FindUserByEmail находит пользователя по email.
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}

	u.Role, err = r.store.Role().Find(roleId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// CheckActive проверяет активность пользователя по id или email.
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
		if errors.Is(err, sql.ErrNoRows) {
			return false, store.ErrRec404
		}
		return false, err
	}

	return u.IsActive, nil
}

// Find находит пользователя по id.
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRec404
		}
		return nil, err
	}

	u.Role, err = r.store.Role().Find(roleId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Delete удаляет пользователя по id.
func (r *UserRepository) Delete(id int64) error {
	_, err := r.store.db.Exec("UPDATE users SET is_active = 0 WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser обновляет данные пользователя.
func (r *UserRepository) Update(ctx context.Context, u *model.User) error {
	current, err := r.Find(u.ID)

	if err != nil {
		return err
	}

	query, values := utils.PrepareUpdateQueryAndValues(current, u)

	if len(values) == 0 {
		return nil
	}

	tx, err := r.store.getTxFromCtx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

// GetUserList возвращает список пользователей с пагинацией.
func (r *UserRepository) GetList(page int64, limit int64) ([]*model.User, error) {
	offset := (page - 1) * limit // Calculate offset for pagination

	rows, err := r.store.db.Query(
		"SELECT id, email FROM users LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr // Assign the error to the named return value
		}
	}()

	var users []*model.User
	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(
			&u.ID,
			&u.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
