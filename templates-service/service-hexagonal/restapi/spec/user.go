package spec

import "time"

type DtoRegisterLoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginInfo struct {
	UserId     string
	Expiration time.Time
}
