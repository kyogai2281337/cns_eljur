package store

import "github.com/kyogai2281337/cns_eljur/pkg/sql/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	CheckActive(interface{}) (bool, error)
	Find(int64) (*model.User, error)
	Delete(int64) error
	Update(*model.User) error
	GetList(page int64, limit int64) ([]*model.User, error)
}

type RoleRepository interface {
	Find(int64) (*model.Role, error)
	CreateRole(string) (*model.Role, error) // CreateRole (setter)
	FindByName(string) (*model.Role, error) // FindRoleByName -> Set RolePermission
	GetList(page int64, limit int64) ([]*model.Role, error)
}

type GroupRepository interface {
	Find(int64) (*model.Group, error)
	Create(*model.Group) (*model.Group, error)
	FindByName(string) (*model.Group, error)
	GetList(page int64, limit int64) ([]*model.Group, error)
	Update(*model.Group) error
}

type CabinetRepository interface {
	Find(int64) (*model.Cabinet, error)
	Create(*model.Cabinet) (*model.Cabinet, error)
	FindByName(string) (*model.Cabinet, error)
	GetList(page int64, limit int64) ([]*model.Cabinet, error)
	Update(*model.Cabinet) error
}

type SubjectRepository interface {
	Find(int64) (*model.Subject, error)
	Create(*model.Subject) (*model.Subject, error)
	FindByName(string) (*model.Subject, error)
	GetList(page int64, limit int64) ([]*model.Subject, error)
	Update(*model.Subject) error
}

type TeacherRepository interface {
	Find(int64) (*model.Teacher, error)
	Create(*model.Teacher) (*model.Teacher, error)
	FindByName(string) (*model.Teacher, error)
	GetList(page int64, limit int64) ([]*model.Teacher, error)
	Update(*model.Teacher) error
}

type SpecializationRepository interface {
	Find(int64) (*model.Specialization, error)
	Create(*model.Specialization) (*model.Specialization, error)
	FindByName(string) (*model.Specialization, error)
	GetList(page int64, limit int64) ([]*model.Specialization, error)
	Update(*model.Specialization) error
}
