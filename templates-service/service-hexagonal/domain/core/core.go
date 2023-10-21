package core

import "time"

type CoreTodo interface {
	PastDeadline(time.Time) (time.Time, bool)
}
