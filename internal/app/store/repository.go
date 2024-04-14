package store

import "github.com/kyogai2281337/cns_eljur/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	CheckActive(interface{}) (bool, error)
	Find(int64) (*model.User, error)
	Delete(int64) error
}
