package models

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserCreate struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserUpdate struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type UserGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
