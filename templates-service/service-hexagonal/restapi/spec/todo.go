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
// 	TodoID string `json:"id"`
// }

// type DtoDeleteTodo struct {
// 	TodoID string `json:"id"`
// }
