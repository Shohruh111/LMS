package models

type Phone struct {
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsFax       bool   `json:"is_fax"`
}

type CreatePhone struct {
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsFax       bool   `json:"is_fax"`
}

type UpdatePhone struct {
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsFax       bool   `json:"is_fax"`
}

type PhonePrimaryKey struct {
	UserId string `json:"user_id"`
}

type GetListPhoneRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListPhoneResponse struct {
	Count  int      `json:"count"`
	Phones []*Phone `json:"phones"`
}
