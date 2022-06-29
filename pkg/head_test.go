package datastruc

import (
	"testing"
)

var testLesser = func(i, j int) bool {
	return i < j
}

func TestHeap(t *testing.T) {
	heap := NewHeap(testLesser, HeapOptionCapacity{100})

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

	if heap.Len() != 0 {
		t.Errorf("expected len of 0 but got %d", heap.Len())
	}

	if !heap.IsEmpty() {
		t.Errorf("expected is empty to be true but was false")
	}

	if got = heap.Peak(); got != 0 {
		t.Errorf("expected peak on empty to return zero but got %d", got)
	}

	if got = heap.Pop(); got != 0 {
		t.Errorf("expected pop on empty to return zero but got %d", got)
	}
}

func TestIsHeapOption(t *testing.T) {
	type test struct {
		name   string
		option HeapOption
	}

	tests := []test{
		{
			name:   "HeapOptionCapacity",
			option: HeapOptionCapacity{100},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.option.isHeapOption()

			if !got {
				t.Errorf("expected a heap option but got false")
			}
		})
	}
}
