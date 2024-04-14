package sqlstore_test

import (
	"github.com/kyogai2281337/cns_eljur/internal/app/model"
	"github.com/kyogai2281337/cns_eljur/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}

func TestRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	email := "nigga@ex.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	var id int64 = 1

	_, err := s.User().Find(id)
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	u, err := s.User().Find(id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u:=model.TestUser(t)
	s.User().Create(u)
	err := s.User().Delete(u.ID)
	assert.NoError(t, err)

}