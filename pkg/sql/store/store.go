package store

import "context"

type Store interface {
	User() UserRepository
	Role() RoleRepository
	GetTables() []string
	Group() GroupRepository
	Cabinet() CabinetRepository
	Subject() SubjectRepository
	Teacher() TeacherRepository
	Specialization() SpecializationRepository
	BeginTx(context.Context) (context.Context, error)
	RollbackTx(context.Context) error
	CommitTx(context.Context) error
	Close() error
}
