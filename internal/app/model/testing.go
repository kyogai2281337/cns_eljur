package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:     "nigga@ex.org",
		Pass:      "password",
		FirstName: "Nigga",
		LastName:  "Pidorasov",
	}
}
