package sqlstore

import (
	"database/sql"

	"github.com/kyogai2281337/cns_eljur/internal/auth/store"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db                   *sql.DB
	userRepository       *UserRepository
	roleRepository       *RoleRepository
	permissionRepository *PermissionRepository
}

func (s *Store) Permission() store.PermissionRepository {
	if s.permissionRepository != nil {
		return s.permissionRepository
	}
	s.roleRepository = &RoleRepository{
		store: s,
	}
	return s.permissionRepository
}

// Constructor
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Role() store.RoleRepository {
	if s.roleRepository != nil {
		return s.roleRepository
	}
	s.roleRepository = &RoleRepository{
		store: s,
	}
	return s.roleRepository
}
