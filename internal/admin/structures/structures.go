package structures

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type GetObjRequest struct {
	TableName string `json:"tname"`
	Id        int64  `json:"id"`
}

type GetUserResponse struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Role      *model.Role `json:"role,omitempty"`
	IsActive  bool        `json:"deleted"`
}

type GetListRequest struct {
	TableName string `json:"tname"`
	Limit     int64  `json:"limit"`
	Page      int64  `json:"page"`
}

type GetListResponse struct {
	Table []interface{} `json:"table"`
}

type GetUserListResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type GetTablesResponse struct {
	Tables []string `json:"tables"`
}

type SetObj struct {
	TableName string      `json:"tablename"`
	Table     interface{} `json:"table"`
}
