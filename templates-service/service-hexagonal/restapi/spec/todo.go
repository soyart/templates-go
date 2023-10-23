package spec

import (
	"time"
)

type DtoCreateTodo struct {
	Text     string    `json:"text"`
	Deadline time.Time `json:"deadline"`
}

// These inputs are taken from URL queries
//
// type DtoMatchTodo struct {
// 	Pattern string `json:"string"`
// }

// type DtoExpireTodo struct {
// 	TodoId string `json:"id"`
// }

// type DtoDeleteTodo struct {
// 	TodoId string `json:"id"`
// }
