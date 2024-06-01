package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        int64         `json:"id,omitempty"`
	Email     string        `json:"email"`
	Pass      string        `json:"password,omitempty"`
	EncPass   string        `json:"-"`
	FirstName string        `json:"first_name,omitempty"`
	LastName  string        `json:"last_name,omitempty"`
	IsActive  bool          `json:"-"`
	Role      *Role         `json:"role,omitempty"`
	PermsSet  *[]Permission `json:"permissions,omitempty"`
}

func (u *User) BeforeCreate() error {
	if err := u.Validate(); err != nil {
		return err
	}
	enc, err := encryptString(u.Pass)
	if err != nil {
		return err
	}

	u.EncPass = enc
	return nil
}

func (u *User) Sanitize() {
	u.Pass = ""
}

func (u *User) ComparePass(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncPass), []byte(password)) == nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Pass, validation.By(RequiredIf(u.EncPass == "")), validation.Length(8, 100)),
	)
}

func encryptString(value string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
