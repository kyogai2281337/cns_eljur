package store

type Store interface {
	User() UserRepository
	Role() RoleRepository
	GetTables() []string
	Group() GroupRepository
	Cabinet() CabinetRepository
	Subject() SubjectRepository
	Teacher() TeacherRepository
	Specialization() SpecializationRepository
}
