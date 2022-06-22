package datastruc

import "sync"

type Heap[T any] interface {
	Len() int
	Peak() T
	Pop() T
	Push(v T)
	PushMany(v ...T)
}

type heap[T any] struct {
	lesser func(i, j T) bool
	data   []T
	lock   *sync.Mutex
}

func NewHeap[T any](lesser func(i, j T) bool) Heap[T] {
	return NewHeapSize[T](lesser, 100)
}

func NewHeapSize[T any](lesser func(i, j T) bool, size int) Heap[T] {
	heap := &heap[T]{
		lesser: lesser,
		data:   make([]T, 0, size),
		lock:   &sync.Mutex{},
	}

	return heap
}

func (heap *heap[T]) Len() int {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	return len(heap.data)
}

func (heap *heap[T]) Peak() T {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	return heap.data[0]
}

func (heap *heap[T]) Pop() T {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	value := heap.data[0]
	lastIndex := len(heap.data) - 1
	heap.data[0] = heap.data[lastIndex]
	heap.data = heap.data[:lastIndex]

	heap.pushDown(0)

	return value
}

func (heap *heap[T]) Push(v T) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	heap.data = append(heap.data, v)
	heap.pushUp(len(heap.data) - 1)
}

func (heap *heap[T]) PushMany(v ...T) {
	heap.data = append(heap.data, v...)
	dataLen := len(heap.data)
	for i := dataLen - 1; i >= dataLen/2; i-- {
		heap.pushUp(i)
	}
}

func (heap *heap[T]) pushUp(index int) {
	for {
		if index <= 0 {
			return
		}

		parentIndex := index / 2

		value := heap.data[index]
		parentValue := heap.data[parentIndex]

		if !heap.lesser(heap.data[index], heap.data[parentIndex]) {
			return
		}

		heap.data[parentIndex] = value
		heap.data[index] = parentValue

		index = parentIndex
	}
}

func (heap *heap[T]) pushDown(index int) {
	for {
		dataLen := len(heap.data)
		if index >= dataLen {
			return
		}

		leftIndex := (index * 2) + 1
		rightIndex := (index * 2) + 2

		var lessIndex int
		var lessValue T

		if leftIndex >= dataLen {
			// is a leaf
			return
		} else if rightIndex >= dataLen {
			// only one child
			lessIndex = leftIndex
			lessValue = heap.data[leftIndex]
		} else {
			// two children
			leftValue := heap.data[leftIndex]
			rightValue := heap.data[rightIndex]
			if heap.lesser(leftValue, rightValue) {
				lessIndex = leftIndex
				lessValue = leftValue
			} else {
				lessIndex = rightIndex
				lessValue = rightValue
			}
		}

		value := heap.data[index]

		if !heap.lesser(lessValue, value) {
			return
		}

		heap.data[index] = lessValue
		heap.data[lessIndex] = value
		index = lessIndex
	}
}
