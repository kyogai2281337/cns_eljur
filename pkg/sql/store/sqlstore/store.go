package sqlstore

import (
	"database/sql"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db                       *sql.DB
	userRepository           *UserRepository
	roleRepository           *RoleRepository
	groupRepository          *GroupRepository
	cabinetRepository        *CabinetRepository
	subjectRepository        *SubjectRepository
	teacherRepository        *TeacherRepository
	specializationRepository *SpecializationRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetTables() []string {
	return []string{"users", "roles", "groups", "cabinets", "subjects", "teachers"}
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

func (s *Store) Group() store.GroupRepository {
	if s.groupRepository != nil {
		return s.groupRepository
	}
	s.groupRepository = &GroupRepository{
		store: s,
	}
	return s.groupRepository
}

func (s *Store) Cabinet() store.CabinetRepository {
	if s.cabinetRepository != nil {
		return s.cabinetRepository
	}
	s.cabinetRepository = &CabinetRepository{
		store: s,
	}
	return s.cabinetRepository
}

func (s *Store) Subject() store.SubjectRepository {
	if s.subjectRepository != nil {
		return s.subjectRepository

	}
	s.subjectRepository = &SubjectRepository{
		store: s,
	}
	return s.subjectRepository
}

func (s *Store) Teacher() store.TeacherRepository {
	if s.teacherRepository != nil {
		return s.teacherRepository
	}
	s.teacherRepository = &TeacherRepository{
		store: s,
	}
	return s.teacherRepository
}

func (s *Store) Specialization() store.SpecializationRepository {
	if s.specializationRepository != nil {
		return s.specializationRepository
	}

	s.specializationRepository = &SpecializationRepository{
		store: s,
	}
	return s.specializationRepository
}
