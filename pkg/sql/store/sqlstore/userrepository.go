package sqlstore

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
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

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
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
		if errors.Is(err, sql.ErrNoRows) {
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
		if errors.Is(err, sql.ErrNoRows) {
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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error in database method SearchPermissions: %s", err)
		}
	}(rows)
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

func (r *UserRepository) UpdateUser(u *model.User) error {
	// Получаем текущую запись пользователя из базы данных
	current, err := r.Find(u.ID)
	if err != nil {
		return err
	}

	// Создаем слайс для хранения значений и строку для SQL-запроса
	var values []interface{}
	query := "UPDATE users SET"

	// Сравниваем поля текущей и новой записи пользователя
	if current.Email != u.Email {
		query += " email = ?,"
		values = append(values, u.Email)
	}
	if current.EncPass != u.EncPass {
		query += " encrypted_password = ?,"
		values = append(values, u.EncPass)
	}
	if current.FirstName != u.FirstName {
		query += " first_name = ?,"
		values = append(values, u.FirstName)
	}
	if current.LastName != u.LastName {
		query += " last_name = ?,"
		values = append(values, u.LastName)
	}
	if current.Role.ID != u.Role.ID {
		query += " role_id = ?,"
		values = append(values, u.Role.ID)
	}

	// Если ни одно поле не было изменено, возвращаем nil
	if len(values) == 0 {
		return nil
	}

	// Добавляем ID пользователя в конец слайса значений
	values = append(values, u.ID)

	// Удаляем последнюю запятую и добавляем условие WHERE в SQL-запрос
	query = query[:len(query)-1] + " WHERE id = ?"

	// Выполняем SQL-запрос
	_, err = r.store.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserList(page int, limit int) (users []*model.User, err error) {
	rows, err := r.store.db.Query(
		"SELECT id, email, encrypted_password, first_name, last_name, role_id FROM users LIMIT ? OFFSET ?",
		limit,
		page*limit,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		u := &model.User{}
		var roleId int64
		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.EncPass,
			&u.FirstName,
			&u.LastName,
			&roleId,
		); err != nil {
			return nil, err
		}
		u.Role, err = r.store.Role().FindRoleById(roleId)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	// Получить права для всех пользователей единоразово
	err = r.SearchPermissionsForUsers(users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) SearchPermissionsForUsers(users []*model.User) error {
	// Создаем слайс для хранения ID пользователей
	var ids []interface{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	// Создаем строку с плейсхолдерами для SQL-запроса
	placeholders := strings.Trim(strings.Repeat("?,", len(users)), ",")

	// Выполняем SQL-запрос
	rows, err := r.store.db.Query(
		"SELECT id_user, id_perm FROM usr_perms WHERE id_user IN ("+placeholders+")",
		ids...,
	)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error in database method SearchPermissionsForUsers: %s", err)
		}
	}(rows)

	// Создаем мапу для хранения разрешений каждого пользователя
	userPermissions := make(map[int64][]model.Permission)

	// Читаем результаты запроса
	for rows.Next() {
		var userId int64
		var permId int32
		if err := rows.Scan(&userId, &permId); err != nil {
			return err
		}

		// Получаем информацию о разрешении
		perm, err := r.store.Permission().FindPermById(permId)
		if err != nil {
			if errors.Is(err, store.ErrRec404) {
				continue
			} else {
				return err
			}
		}

		// Добавляем разрешение в мапу
		userPermissions[userId] = append(userPermissions[userId], *perm)
	}

	// Присваиваем разрешения пользователям
	for _, user := range users {
		perms := userPermissions[user.ID]
		userPerms := make([]model.Permission, len(perms))
		copy(userPerms, perms)
		user.PermsSet = &userPerms
	}

	return nil
}
