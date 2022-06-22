package datastruc

import (
	"testing"
)

func TestHeap(t *testing.T) {
	testLesser := func(i, j int) bool {
		return i < j
	}

	heap := NewHeap[int](testLesser)

	heap.PushMany(10, 1, 8)
	heap.Push(5)
	heap.Push(7)
	heap.Push(3)
	heap.Push(2)
	heap.Push(2)

	expectResults := []int{1, 2, 2, 3, 5, 7, 8, 10}
	var got int
	for _, expect := range expectResults {
		if got = heap.Peak(); got != expect {
			t.Errorf("expected %d but got %d", expect, got)
		}
		if got = heap.Pop(); got != expect {
			t.Errorf("expected %d but got %d", expect, got)
		}
	}
}
