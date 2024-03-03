package models

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"data"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckEmail struct {
	RequestId string `json:"request_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type ConfirmCode struct {
	RequestId string `json:"request_id"`
}

type CheckCode struct {
	Code      string `json:"verify_code"`
	RequestID string `json:"request_id"`
}

type UpdatePassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
