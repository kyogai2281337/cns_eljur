package store

import "github.com/kyogai2281337/cns_eljur/internal/auth/model"

// UserRepository only auth!
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	CheckActive(interface{}) (bool, error)
	Find(int64) (*model.User, error)
	Delete(int64) error
	SearchPermissions(*model.User) (error, *[]model.Permission)
}

type RoleRepository interface {
	// Find Нужно добавить методы, в первую очередь для работы БД. То есть:
	// 1. Find(getter)
	Find(int64) (*model.Role, error)
	// Create (setter)
	Create(string) (*model.Role, error)
}

type PermissionRepository interface {
}
