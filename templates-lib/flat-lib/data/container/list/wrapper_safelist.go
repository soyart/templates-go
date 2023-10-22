package list

import "sync"

// SafeListWrapper[T] wraps BasicList[T] and uses sync.RWMutex to avoid data races.
// L was added to make sure that the underlying list type is always accessible from the instance type,
// for example, SafeListWrapper[float64, *Queue[float64]]
// if L was not the type parameter, then a safe uint8 stack, safe uint8 queue, etc,
// will be indifferentiable with the same type `SafeListWrapper[uint8]`.
// SafeListWrapper[T, BasicList[T]] also implements BasicList[T]
type SafeListWrapper[T any, L BasicList[T]] struct {
	mut       *sync.RWMutex
	basicList L
}

// WrapSafeList[T] wraps a BasicList[T] into SafeListWrapper[T],
// where T is the underlying entity (item) type and L is the underlying BasicList[T] type.
// If you're wrapping a variable `foo“ of type `*Stack[uint8]`, then call this function with:
// WrapSafeList[uint8, *Stack[uint8]](foo)
func WrapSafeList[T any, L BasicList[T]](basicList L) *SafeListWrapper[T, L] {
	return &SafeListWrapper[T, L]{
		basicList: basicList,
		mut:       &sync.RWMutex{},
	}
}

func (w *SafeListWrapper[T, L]) Push(x T) {
	w.mut.Lock()
	defer w.mut.Unlock()

	w.basicList.Push(x)
}

func (w *SafeListWrapper[T, L]) PushSlice(x []T) {
	w.mut.Lock()
	defer w.mut.Unlock()

	w.basicList.PushSlice(x)
}

func (w *SafeListWrapper[T, L]) Pop() *T {
	w.mut.Lock()
	defer w.mut.Unlock()

	return w.basicList.Pop()
}

func (w *SafeListWrapper[T, L]) Len() int {
	w.mut.RLock()
	defer w.mut.RUnlock()

	return w.basicList.Len()
}

func (w *SafeListWrapper[T, L]) IsEmpty() bool {
	w.mut.RLock()
	defer w.mut.RUnlock()

	return w.basicList.IsEmpty()
}

func (w *SafeListWrapper[T, L]) IsSafe() bool {
	return true
}
