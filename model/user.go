package model

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Auth struct {
	Token  string `json:"token"`
	Role   string `json:"role"`
	UserId int    `json:"userid"`
}
