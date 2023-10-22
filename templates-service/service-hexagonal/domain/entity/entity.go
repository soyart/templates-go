package entity

import "time"

// HexArch allows reusable entity structure
// to be used in business domains and database storage.
type Todo struct {
	Id        string    `json:"id" gorm:"primaryKey;column:id"`
	Text      string    `json:"text" gorm:"column:text"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	Deadline  time.Time `json:"deadline" gorm:"column:deadline"`
}

type User struct {
	Id       string `json:"id" gorm:"primaryKey;column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"-" gorm:"column:password"`
}
