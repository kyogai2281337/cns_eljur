package adminPanel_structures

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type GetObjRequest struct {
	TableName string `json:"tname"`
	Id        int64  `json:"id"`
}

type GetObjResponse struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Role      *model.Role `json:"role,omitempty"`
	IsActive  bool        `json:"isActive"`
}

type GetListRequest struct {
	TableName string `json:"tname"`
	Limit     int64  `json:"limit"`
	Page      int64  `json:"page"`
}

type GetListResponse struct {
	Table []TableStruct `json:"table"`
}
type TableStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
