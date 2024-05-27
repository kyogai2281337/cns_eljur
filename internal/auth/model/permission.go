package model

type Permission struct {
	Id      int32 `json:"id"`
	IsAdmin bool  `json:"isAdmin"`
}
