package store

import "github.com/kyogai2281337/cns_eljur/internal/auth/model"

// only auth!
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	CheckActive(interface{}) (bool, error)
	Find(int64) (*model.User, error)
	Delete(int64) error
}

type RoleRepository interface {
	// Нужно добавить методы, в первую очередь для работы БЛ. То есть:
	// 1. Find(getter)
	Find(int64) (*model.Role, error)
	// 2. Create(setter)
	Create(string) (*model.Role, error)
}

type PermissionRepository interface {
	CheckRole(*model.Role) (bool, error)
}
