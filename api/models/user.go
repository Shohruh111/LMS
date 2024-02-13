package models

type LoginUser struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

type CheckEmail struct {
	Email string `json:"email"`
}

type ConfirmCode struct {
	RequestId string `json:"request_id"`
}

type CheckCode struct {
	Code      int    `json:"verify_code"`
	RequestID string `json:"request_id"`
}

type UpdatePassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id          string `json:"id"`
	RoleId      string `json:"role_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserCreate struct {
	RoleId      string `json:"role_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserUpdate struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type UserGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
