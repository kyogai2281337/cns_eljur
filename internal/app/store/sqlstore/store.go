package sqlstore

import (
	"database/sql"
	"github.com/kyogai2281337/cns_eljur/internal/app/store"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
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
