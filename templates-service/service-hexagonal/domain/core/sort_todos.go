package core

import (
	"sort"

	"example.com/servicehex/domain/entity"
)

type EnumSortMode = int

const (
	SortInvalid EnumSortMode = iota
	SortNoSort
	SortDeadlineDescending
	SortDeadlineAscending
	SortTextDescending
	SortTextAscending

	DefaultSort = SortDeadlineDescending
)

// SortTodos sorts `todos` *in-place*
func SortTodos(todos []entity.Todo, mode EnumSortMode) {
	switch mode {
	case SortInvalid:
		// Should not happen, unless we want to
		// treat it as SortDeadlineDescending
		panic("got invalid eum SortInvalid")

	case SortNoSort:
		return

	case SortDeadlineDescending:
		sortTodoByDeadlineDescending(todos)

	case SortDeadlineAscending:
		sortTodoByDeadlineAscending(todos)

	case SortTextDescending:
		sortTodoByTextDescending(todos)

	case SortTextAscending:
		sortTodoByTextAscending(todos)
	}
}

func sortTodoByTextDescending(todos []entity.Todo) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Text < todos[j].Text
	})
}

func sortTodoByTextAscending(todos []entity.Todo) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Text > todos[j].Text
	})
}

func sortTodoByDeadlineDescending(todos []entity.Todo) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Deadline.Before(todos[j].Deadline)
	})
}

func sortTodoByDeadlineAscending(todos []entity.Todo) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Deadline.After(todos[j].Deadline)
	})
}
