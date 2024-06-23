package store

type Store interface {
	User() UserRepository
	Role() RoleRepository
	GetTables() []string
}
