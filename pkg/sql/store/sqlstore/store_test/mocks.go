package store_test

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(u *model.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) CheckActive(param interface{}) (bool, error) {
	args := m.Called(param)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) Find(id int64) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserList(page int64, limit int64) ([]*model.User, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]*model.User), args.Error(1)
}
