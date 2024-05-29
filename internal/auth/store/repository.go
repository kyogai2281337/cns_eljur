package store

import "github.com/kyogai2281337/cns_eljur/internal/auth/model"

// UserRepository only auth!
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	CheckActive(interface{}) (bool, error)
	Find(int64) (*model.User, error)
	Delete(int64) error
}

type RoleRepository interface {
	// FindRoleById Find Нужно добавить методы, в первую очередь для работы БД. То есть:
	// 1. Find(getter)
	FindRoleById(int64) (*model.Role, error)
	// CreateRole (setter)
	CreateRole(string) (*model.Role, error)
	// FindRoleByName -> Set RolePermission
	FindRoleByName(string) (*model.Role, error)
}

type PermissionRepository interface {
}
