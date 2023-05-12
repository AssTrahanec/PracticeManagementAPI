package model

type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Role  string `json:"role"`
}

type UserResponse struct {
	User
	Password string `json:"-"`
}
