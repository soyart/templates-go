package list

import (
	"container/heap"
	"math/big"
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"

	"example.com/libflat/data"
)

type foo[T constraints.Ordered] struct {
	name  string
	value T
}

// Implements data.Valuer[T]
func (f foo[T]) GetValue() T {
	return f.value
}

func TestPq(t *testing.T) {
	highest := foo[int]{name: "b", value: 100}
	lowest := foo[int]{name: "d", value: 0}

	items := []foo[int]{
		{name: "a", value: 69},
		highest,
		{name: "c", value: 12},
		lowest,
	}

	minHeapResults := testPop(t, MinHeap, items)
	for _, minHeapResult := range minHeapResults {
		if minHeapResult != lowest {
			t.Fatalf("unexpected MinHeap results - expected %+v, got %+v\n", lowest, minHeapResult)
		}
	}

	maxHeapResults := testPop(t, MaxHeap, items)
	for _, maxHeapResult := range maxHeapResults {
		if maxHeapResult != highest {
			t.Fatalf("unexpected MaxHeap results - expected %+v, got %+v\n", highest, maxHeapResult)
		}
	}

	testArbitaryUpdate(t)
}

func testPop[T constraints.Ordered](t *testing.T, pqType HeapType, items []foo[T]) []foo[T] {
	pq := NewPriorityQueue[T](pqType)
	pqCustom := NewPriorityQueueCustom(pqType, lessOrdered[T])
	queues := []*PriorityQueue[T]{pq, pqCustom}

	var ret []foo[T]
	for _, q := range queues {
		for _, item := range items {
			heap.Push(q, item)
		}

		p := heap.Pop(q)
		popped, ok := p.(foo[T])
		if !ok {
			t.Fatalf("type assertion to *foo failed, got type: %s", reflect.TypeOf(p))
		}
		ret = append(ret, popped)
	}
	return ret
}

func testArbitaryUpdate(t *testing.T) {
	// Test with Valuer[float64]
	hundred := foo[float64]{name: "hundred", value: 100}
	seventy := foo[float64]{name: "seventy", value: 70}
	zero := foo[float64]{name: "zero", value: 0}
	foosFloat := []foo[float64]{
		{name: "a", value: 69},
		hundred,
		{name: "b", value: 71},
		zero,
		seventy,
	}

	// Arbitary pushes and inits
	pq := NewPriorityQueue[float64](MaxHeap)
	pqCustom := NewPriorityQueueCustom(MaxHeap, lessOrdered[float64])
	queues := []*PriorityQueue[float64]{pq, pqCustom}

	for _, q := range queues {
		for _, item := range foosFloat {
			heap.Push(q, item)
		}

		p := heap.Pop(q)
		popped, ok := p.(foo[float64])
		if !ok {
			t.Fatalf("type assertion to *foo failed, got type: %s", reflect.TypeOf(p))
		}
		if popped != hundred {
			t.Fatalf("unexpected MaxHeap results - expected %+v, got %+v\n", hundred, popped)
		}

		q.HeapType = MinHeap
		heap.Init(q)
		p = heap.Pop(q)
		popped, ok = p.(foo[float64])
		if !ok {
			t.Fatalf("type assertion to *foo failed, got type: %s", reflect.TypeOf(p))
		}
		if popped != zero {
			t.Fatalf("unexpected MinHeap results - expected %+v, got %+v\n", zero, popped)
		}
	}
}

type bar struct {
	val *big.Int
}

func (b *bar) GetValue() *big.Int {
	return b.val
}

func TestPQCmp(t *testing.T) {
	a := &bar{val: big.NewInt(69)}
	b := &bar{val: big.NewInt(70)}
	c := &bar{val: big.NewInt(100)}
	d := &bar{val: big.NewInt(1000000)}

	t.Run("MaxHeap with Cmp", func(t *testing.T) {
		testPqCmpMax(t, []*bar{a, d, c, b}, d)
	})
	t.Run("MinHeap with Cmp", func(t *testing.T) {
		testPqCmpMin(t, []*bar{a, d, c, b}, a)
	})

	lol(a) // Compiles and no panic
}

func lol(item data.GetValuer[*big.Int]) {

}

func testPqCmpMax(t *testing.T, messy []*bar, max *bar) {
	maxPq := NewPriorityQueueCmp[*big.Int](MaxHeap)
	maxPqCustom := NewPriorityQueueCustom(MaxHeap, lessCmp[*big.Int])
	queues := []*PriorityQueue[*big.Int]{maxPq, maxPqCustom}

	for _, q := range queues {
		for _, item := range messy {
			heap.Push(q, item)
		}
		if popped := heap.Pop(q); popped != nil {
			actual := popped.(*bar)
			if actual != max {
				t.Fatalf("unexpected max heap result: expected %v, got %v\n", max.GetValue(), actual.GetValue())
			}
		}
	}
}

func testPqCmpMin(t *testing.T, messy []*bar, min *bar) {
	minPq := NewPriorityQueueCmp[*big.Int](MinHeap)
	minPqCustom := NewPriorityQueueCustom(MinHeap, lessCmp[*big.Int])
	queues := []*PriorityQueue[*big.Int]{minPq, minPqCustom}

	for _, q := range queues {
		for _, item := range messy {
			heap.Push(q, item)
		}
		if popped := heap.Pop(q); popped != nil {
			actual := popped.(*bar)
			if actual != min {
				t.Fatalf("unexpected min heap result: expected %v, got %v\n", min.GetValue(), actual.GetValue())
			}
		}
	}
}
