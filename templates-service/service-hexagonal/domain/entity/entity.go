package entity

import "time"

type Todo struct {
	Id        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Deadline  time.Time `json:"deadline"`
}
