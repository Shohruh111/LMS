package models

type Role struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RoleCreate struct {
	Type string `json:"type"`
}

type RoleUpdate struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type RolePrimaryKey struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type RoleGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type RoleGetListResponse struct {
	Count int     `json:"count"`
	Roles []*Role `json:"roles"`
}
