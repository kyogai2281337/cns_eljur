package sqlstore

import (
	"database/sql"
	"errors"
	"log"
	"strings"

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

func (r *RoleRepository) SearchPermissionsForRoles(roles []*model.Role) error {
	// Создаем слайс для хранения ID ролей
	var ids []interface{}
	for _, role := range roles {
		ids = append(ids, role.ID)
	}

	// Создаем строку с плейсхолдерами для SQL-запроса
	placeholders := strings.Trim(strings.Repeat("?,", len(roles)), ",")

	// Выполняем SQL-запрос
	rows, err := r.store.db.Query(
		"SELECT id_role, id_perm FROM role_perms WHERE id_role IN ("+placeholders+")",
		ids...,
	)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error in database method SearchPermissionsForRoles: %s", err)
		}
	}(rows)

	// Создаем мапу для хранения разрешений каждой роли
	rolePermissions := make(map[int64][]model.Permission)

	// Читаем результаты запроса
	for rows.Next() {
		var roleId int32
		var permId int32
		if err := rows.Scan(&roleId, &permId); err != nil {
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
		rolePermissions[int64(roleId)] = append(rolePermissions[int64(roleId)], *perm)
	}

	// Присваиваем разрешения ролям

	for _, role := range roles {
		perms := rolePermissions[int64(role.ID)]
		rolePerms := make([]model.Permission, len(perms))
		copy(rolePerms, perms)
		role.PermsSet = &rolePerms
	}

	return nil
}

func (r *RoleRepository) UpdateRole(role *model.Role) error {
	// Получаем текущую запись роли из базы данных
	current, err := r.FindRoleById(int64(role.ID))
	if err != nil {
		return err
	}

	// Создаем слайс для хранения значений и строку для SQL-запроса
	var values []interface{}
	query := "UPDATE roles SET"

	// Сравниваем поля текущей и новой записи роли
	if current.Name != role.Name {
		query += " name = ?,"
		values = append(values, role.Name)
	}

	// Если ни одно поле не было изменено, возвращаем nil
	if len(values) == 0 {
		return nil
	}

	// Добавляем ID роли в конец слайса значений
	values = append(values, role.ID)

	// Удаляем последнюю запятую и добавляем условие WHERE в SQL-запрос
	query = query[:len(query)-1] + " WHERE id = ?"

	// Выполняем SQL-запрос
	_, err = r.store.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) GetRoleList(page int, limit int) (roles []*model.Role, err error) {
	rows, err := r.store.db.Query(
		"SELECT id, name FROM roles LIMIT ? OFFSET ?",
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
		role := &model.Role{}
		if err := rows.Scan(
			&role.ID,
			&role.Name,
		); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	// Получить права для всех ролей единоразово
	err = r.SearchPermissionsForRoles(roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
