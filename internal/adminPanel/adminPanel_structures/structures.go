package adminPanel_structures

type GetObjRequest struct {
	TName string `json:"tname"`
	Id    int64  `json:"id"`
}

type GetObjResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetListRequest struct {
	TName string `json:"tname"`
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
}

type GetListResponse struct {
	Table []TableStruct `json:"table"`
}
type TableStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
