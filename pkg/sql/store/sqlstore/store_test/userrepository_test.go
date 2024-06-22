package store_test_test

import (
	"testing"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore/store_test"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserIntegration(t *testing.T) {
	mockRepo := new(store_test.MockUserRepository)

	user := &model.User{ID: 1, Email: "levandr@example.com"}

	mockRepo.On("UpdateUser", user).Return(nil)
	mockRepo.On("Find", user.ID).Return(user, nil)

	err := mockRepo.UpdateUser(user)
	assert.Nil(t, err)

	updatedUser, err := mockRepo.Find(user.ID)
	assert.Nil(t, err)

	assert.Equal(t, "levandr@example.com", updatedUser.Email)

	mockRepo.AssertExpectations(t)
}
