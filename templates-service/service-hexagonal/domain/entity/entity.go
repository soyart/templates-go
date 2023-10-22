package entity

import "time"

// HexArch allows reusable entity structure
// to be used in business domains and database storage.
type Todo struct {
	Id       string    `json:"id" gorm:"primaryKey;column:id"`
	UserId   string    `json:"user_id" gorm:"column:user_id;notnull"`
	Text     string    `json:"text" gorm:"column:text;notnull"`
	Deadline time.Time `json:"deadline" gorm:"column:deadline"`

	Expired   bool      `json:"expired" gorm:"column:expired;notnull"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull"`
}

type User struct {
	Id       string `json:"id" gorm:"primaryKey;column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"-" gorm:"column:password"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull"`
}
