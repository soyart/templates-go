package domain

import (
	"github.com/google/uuid"
	"time"
)

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

func (s Status) String() string {
	switch s {
	case StatusTodo:
		return "TODO"
	case StatusInProgress:
		return "IN_PROGRESS"
	case StatusDone:
		return "DONE"
	default:
		return ""
	}
}

type TodoCreateInput struct {
	Title       string    `validate:"required"`
	Status      Status    `validate:"required"`
	Description string    `validate:"required"`
	TodoDate    time.Time `validate:"required"`
}

func (t TodoCreateInput) ToEntity() Todo {
	return Todo{
		ID:          uuid.NewString(),
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		TodoDate:    t.TodoDate,
		CreatedAt:   time.Now(),
	}
}

type Todo struct {
	ID          string
	Title       string
	Status      Status
	Description string
	TodoDate    time.Time
	CreatedAt   time.Time
}

func NewTodo(
	id string,
	title string,
	description string,
	status Status,
	todoDate time.Time,
	createdAt time.Time,
) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		TodoDate:    todoDate,
		CreatedAt:   createdAt,
	}
}
