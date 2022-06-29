package datastruc

import "sync"

type SyncHeap[T any] interface {
	Heap[T]
	PopIf(shouldPop func(v T) bool) (popped T, wasPopped bool)
}

type syncHeap[T any] struct {
	backingHeap Heap[T]
	lock        *sync.Mutex
}

func NewSyncHeap[T any](
	lesser func(i, j T) bool,
	options ...HeapOption,
) SyncHeap[T] {
	backingHeap := NewHeap[T](lesser, options...)

	return &syncHeap[T]{
		backingHeap: backingHeap,
		lock:        &sync.Mutex{},
	}
}

func (heap *syncHeap[T]) Clear() {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	heap.backingHeap.Clear()
}

func (heap *syncHeap[T]) IsEmpty() bool {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	return heap.backingHeap.IsEmpty()
}

func (heap *syncHeap[T]) Len() int {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	return heap.backingHeap.Len()
}

func (heap *syncHeap[T]) Peak() (least T) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	return heap.backingHeap.Peak()
}

func (heap *syncHeap[T]) Pop() (least T) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	return heap.backingHeap.Pop()
}

func (heap *syncHeap[T]) PopIf(
	shouldPop func(v T) bool,
) (popped T, wasPopped bool) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	backingHeap := heap.backingHeap

	if backingHeap.IsEmpty() {
		return popped, false
	}

	least := backingHeap.Peak()

	if !shouldPop(least) {
		return popped, false
	}

	return backingHeap.Pop(), true
}

func (heap *syncHeap[T]) Push(v T) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	heap.backingHeap.Push(v)
}

func (heap *syncHeap[T]) PushMany(v ...T) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	heap.backingHeap.PushMany(v...)
}
