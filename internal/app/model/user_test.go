package model_test

import (
	"github.com/kyogai2281337/cns_eljur/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{name: "valid", u: func() *model.User { return model.TestUser(t) }, isValid: true},
		{name: "nilEmail", u: func() *model.User { u := model.TestUser(t); u.Email = ""; return u }, isValid: false},
		{name: "errEmail", u: func() *model.User { u := model.TestUser(t); u.Email = "negrila2281337"; return u }, isValid: false},
		{name: "encNotNil", u: func() *model.User { u := model.TestUser(t); u.Pass = ""; u.EncPass = "negrila"; return u }, isValid: true},
		{name: "nilPass", u: func() *model.User { u := model.TestUser(t); u.Pass = ""; return u }, isValid: false},
	}
	for _, el := range testCases {
		t.Run(el.name, func(t *testing.T) {
			if el.isValid {
				assert.NoError(t, el.u().Validate())
			} else {
				assert.Error(t, el.u().Validate())
			}
		})
	}
}

func TestUserBeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotNil(t, u.EncPass)
}
