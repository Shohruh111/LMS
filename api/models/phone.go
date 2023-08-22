package models

type Phone struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"descriprion"`
	IsFax       bool   `json:"is_fax"`
}

type CreatePhone struct {
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"descriprion"`
	IsFax       bool   `json:"is_fax"`
}

type UpdatePhone struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"descriprion"`
	IsFax       bool   `json:"is_fax"`
}

type PhonePrimaryKey struct {
	Id     string `json:"id"`
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
