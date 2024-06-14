package store

import "github.com/kyogai2281337/cns_eljur/pkg/sql/model"

type UserRepository interface {
	Create(*model.User) error
	FindUserByEmail(string) (*model.User, error)
	CheckActive(interface{}) (bool, error)
	Find(int64) (*model.User, error)
	Delete(int64) error
	UpdateUser(*model.User) error
	GetUserList(page int64, limit int64) ([]*model.User, error)
	SearchPermissionsForUsers(users []*model.User) error
}

type RoleRepository interface {
	FindRoleById(int64) (*model.Role, error)
	CreateRole(string) (*model.Role, error)     // CreateRole (setter)
	FindRoleByName(string) (*model.Role, error) // FindRoleByName -> Set RolePermission
	GetRoleList(page int64, limit int64) ([]*model.Role, error)
	SearchPermissionsForRoles(roles []*model.Role) error
}

type PermissionRepository interface {
	FindPermById(int32) (*model.Permission, error)
}
