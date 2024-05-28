package model

type Permission struct {
	Id       int32  `json:"id"`
	Name     string `json:"isAdmin"`
	Endpoint string `json:"endpoint"`
}
