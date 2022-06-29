package datastruc

import "testing"

func TestSyncHeap(t *testing.T) {
	heap := NewSyncHeap(testLesser, HeapOptionCapacity{100})

	heap.PushMany(76, 23, 77, 10)
	heap.Push(22)
	heap.Push(22)
	heap.Push(88)
	heap.Push(99)
	heap.PushMany(-100, 100)

	expectResults := []int{-100, 10, 22, 22, 23, 76, 77, 88, 99, 100}

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

func TestPopIf(t *testing.T) {
	type test struct {
		name            string
		values          []int
		value           int
		expectPopped    int
		expectWasPopped bool
	}

	tests := []test{
		{
			name:            "with popped",
			values:          []int{7, 6, 5},
			value:           4,
			expectPopped:    5,
			expectWasPopped: true,
		},
		{
			name:            "without popped",
			values:          []int{7, 6, 5},
			value:           8,
			expectPopped:    0,
			expectWasPopped: false,
		},
		{
			name:            "with empty",
			values:          []int{},
			value:           -1,
			expectPopped:    0,
			expectWasPopped: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			heap := NewSyncHeap(testLesser)
			heap.PushMany(tc.values...)

			gotPopped, gotWasPopped := heap.PopIf(func(v int) bool {
				return tc.value < v
			})

			if gotWasPopped != tc.expectWasPopped {
				t.Fatalf(
					"expected was popped to be %t but got %t",
					tc.expectWasPopped,
					gotWasPopped,
				)
			}

			if gotPopped != tc.expectPopped {
				t.Errorf(
					"expected popped to be %d but was %d",
					tc.expectPopped,
					gotPopped,
				)
			}
		})
	}
}
