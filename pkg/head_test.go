package datastruc

import (
	"testing"
)

var testLesser = func(i, j int) bool {
	return i < j
}

func TestHeap(t *testing.T) {
	heap := NewHeap(testLesser)

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

func TestPopIfLeast(t *testing.T) {
	type test struct {
		name           string
		values         []int
		value          int
		expectPopped   int
		expectWasLeast bool
	}

	tests := []test{
		{
			name:           "with least",
			values:         []int{2, 3, 4},
			value:          1,
			expectPopped:   2,
			expectWasLeast: true,
		},
		{
			name:           "without least",
			values:         []int{2, 3, 4},
			value:          5,
			expectPopped:   0,
			expectWasLeast: false,
		},
		{
			name:           "with empty",
			values:         []int{},
			value:          -1,
			expectPopped:   0,
			expectWasLeast: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			heap := NewHeap(testLesser)
			heap.PushMany(tc.values...)
			gotPopped, gotWasLeast := heap.PopIfLeast(tc.value)

			if gotWasLeast != tc.expectWasLeast {
				t.Errorf(
					"expect wasLeast to be %t but got %t",
					tc.expectWasLeast,
					gotWasLeast,
				)
			}

			if gotPopped != tc.expectPopped {
				t.Errorf(
					"expect popped to be %d but got %d",
					tc.expectPopped,
					gotPopped,
				)
			}
		})
	}
}

func TestPopIfNotLeast(t *testing.T) {
	type test struct {
		name              string
		values            []int
		value             int
		expectPopped      int
		expectWasNotLeast bool
	}

	tests := []test{
		{
			name:              "with not least",
			values:            []int{2, 3, 4},
			value:             3,
			expectPopped:      2,
			expectWasNotLeast: true,
		},
		{
			name:              "without not least",
			values:            []int{2, 3, 4},
			value:             1,
			expectPopped:      0,
			expectWasNotLeast: false,
		},
		{
			name:              "with empty",
			values:            []int{},
			value:             1,
			expectPopped:      0,
			expectWasNotLeast: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			heap := NewHeap(testLesser)
			heap.PushMany(tc.values...)
			gotPopped, gotWasNotLeast := heap.PopIfNotLeast(tc.value)

			if gotWasNotLeast != tc.expectWasNotLeast {
				t.Errorf(
					"expect wasNotLeast to be %t but got %t",
					tc.expectWasNotLeast,
					gotWasNotLeast,
				)
			}

			if gotPopped != tc.expectPopped {
				t.Errorf(
					"expect popped to be %d but got %d",
					tc.expectPopped,
					gotPopped,
				)
			}
		})
	}
}
