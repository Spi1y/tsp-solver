package tasks

import (
	"container/heap"
	"fmt"
)

// heapRecord is a record of the heap
type heapRecord struct {
	task *Task
	// The projected path legth of the task
	distance int

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

// A Heap implements heap.Interface and holds HeapRecord.
type Heap struct {
	slice []*heapRecord

	trimValue int
}

// Len returns len of the heap
func (h *Heap) Len() int { return len(h.slice) }

// Less is a comparison function, required for heap.interface
func (h *Heap) Less(i, j int) bool {
	// We have a minHeap, which means top record is a record with a lowest distance
	return h.slice[i].distance < h.slice[j].distance
}

// Swap swaps elements, required for heap.interface
func (h *Heap) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
	h.slice[i].index = i
	h.slice[j].index = j
}

// Push pushes a new element to the heap, required for heap.interface
func (h *Heap) Push(x interface{}) {
	n := len(h.slice)
	item := x.(*heapRecord)
	item.index = n
	h.slice = append(h.slice, item)
}

// Pop pops a top element from the heap deleting it. Required for heap.interface
func (h *Heap) Pop() interface{} {
	old := h.slice
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	h.slice = old[0 : n-1]
	return item
}

// NewHeapQueue creates and returns new heap queue
func NewHeapQueue() *Heap {
	h := &Heap{
		slice:     nil,
		trimValue: -1,
	}

	heap.Init(h)

	return h
}

// Insert inserts several records to the queue
func (h *Heap) Insert(tasks []*Task) {
	// A quick path for an empty insertion
	if len(tasks) == 0 {
		return
	}

	for _, task := range tasks {
		heap.Push(h, task)
	}
}

// TrimTail should trim records from the tail of the queue with a distance greater
// than the given argument. However, it is very costly to do so with the heap, so for
// now, we remember trimming value and use it in PopFirst and IsEmpty to determine
// if the queue should be empty.
func (h *Heap) TrimTail(distance int) {
	if (h.trimValue == -1) || (h.trimValue > distance) {
		h.trimValue = distance
	}
}

// IsEmpty checks if there is no records in the list.
func (h *Heap) IsEmpty() bool {
	if len(h.slice) == 0 {
		return true
	}

	if h.slice[0].distance >= h.trimValue {
		return true
	}

	return false
}

// PopFirst gets the task from the first record in the list and
// removes it from the list.
// If list is empty, it returns nil.
func (h *Heap) PopFirst() *Task {
	if h.IsEmpty() {
		return nil
	}

	if h.slice[0].distance >= h.trimValue {
		return nil
	}

	return heap.Pop(h).(heapRecord).task
}

// String implements the Stringer interface
// Used mainly for testing
func (h *Heap) String() string {
	return fmt.Sprintf("tasks.Heap:%v", h.slice)
}
