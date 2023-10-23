package spec

import "time"

type DtoRegisterLoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginInfo struct {
	UserID     string
	Expiration time.Time
}
