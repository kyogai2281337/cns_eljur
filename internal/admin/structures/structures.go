package structures

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type GetObjRequest struct {
	TableName string `json:"tablename"`
	Id        int64  `json:"id"`
}

type GetRoleResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
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
	TableName string `json:"tablename"`
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

type GetGroupResponse struct {
	ID             int64                 `bson:"_id,omitempty" json:"id"`
	Specialization *model.Specialization `json:"specialization"`
	Name           string                `json:"name"`
	MaxPairs       int                   `json:"max_pairs"`
	//SpecPlan       map[*model.Subject]int `json:"-"`
}

type GetSpecializationResponse struct {
	ID     int64  `bson:"_id,omitempty" json:"id"`
	Name   string `json:"name"`
	Course int    `json:"course"`
	//EduPlan map[*model.Subject]int `json:"-"`
	PlanId    string        `json:"plan_id"`
	ShortPlan map[int64]int `json:"short_plan"`
}

type GetCabinetResponse struct {
	ID   int64         `bson:"_id,omitempty" json:"id"`
	Name string        `json:"name"`
	Type model.CabType `json:"type,omitempty"`
}
type GetSubjectResponse struct {
	ID               int64         `bson:"_id,omitempty" json:"id"`
	Name             string        `json:"name"`
	RecommendCabType model.CabType `json:"type"`
}
type GetTeacherResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	//Links            map[*model.Group][]*model.Subject `json:"full_links"` // todo impl
	LinksID          string            `bson:"_id" json:"links_id,omitempty"`
	RecommendSchCap_ int               `json:"capacity"`
	Sl               map[int64][]int64 `json:"links,omitempty"`
}

type CreateRequest struct {
	Table string      `json:"tablename"`
	Data  interface{} `json:"data"`
}

type CreateTeacherRequest struct {
	ID               int64                             `bson:"_id,omitempty" json:"id"`
	Name             string                            `json:"name"`
	Links            map[*model.Group][]*model.Subject `json:"type"`
	LinksID          int64                             `json:"links_id"`
	RecommendSchCap_ int                               `json:"capacity"`
	SL               map[int64][]int64                 `json:"links"`
}
