package tasks

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/Spi1y/tsp-solver/solver2/types"
)

// Queue implements tasks queue based on binary heap
type Queue struct {
	slice []*heapRecord

	trimValue int
}

// heapRecord is a record of the heap
type heapRecord struct {
	task Task

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

// Len returns len of the heap
func (h *Queue) Len() int { return len(h.slice) }

// Less is a comparison function, required for heap.interface
func (h *Queue) Less(i, j int) bool {
	// We have a minHeap, which means top record is a record with a lowest distance
	return h.slice[i].task.Estimate < h.slice[j].task.Estimate
}

// Swap swaps elements, required for heap.interface
func (h *Queue) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
	h.slice[i].index = i
	h.slice[j].index = j
}

// Push pushes a new element to the heap, required for heap.interface
func (h *Queue) Push(x interface{}) {
	n := len(h.slice)
	item := x.(*heapRecord)
	item.index = n
	h.slice = append(h.slice, item)
}

// Pop pops a top element from the heap deleting it. Required for heap.interface
func (h *Queue) Pop() interface{} {
	old := h.slice
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	h.slice = old[0 : n-1]
	return item
}

// NewHeapQueue creates and returns new heap queue
func NewHeapQueue() *Queue {
	h := &Queue{
		slice:     nil,
		trimValue: -1,
	}

	return h
}

// Insert inserts several records to the queue
func (h *Queue) Insert(tasks []Task) {
	// A quick path for an empty insertion
	if len(tasks) == 0 {
		return
	}

	for _, task := range tasks {
		rec := &heapRecord{
			task: task,
		}
		heap.Push(h, rec)
	}
}

// TrimTail should trim records from the tail of the queue with a distance greater
// than the given argument. However, it is very costly to do so with the heap, so for
// now, we remember the trimming value and use it in PopFirst and IsEmpty to determine
// if the queue should be empty.
func (h *Queue) TrimTail(distance types.Distance) {
	if (h.trimValue == -1) || (h.trimValue > int(distance)) {
		h.trimValue = int(distance)
	}
}

// IsEmpty checks if there is no records in the list.
func (h *Queue) IsEmpty() bool {
	if len(h.slice) == 0 {
		return true
	}

	if (h.trimValue != -1) && (int(h.slice[0].task.Estimate) >= h.trimValue) {
		return true
	}

	return false
}

// PopFirst gets the task from the first record in the list and
// removes it from the list.
// If list is empty, it returns nil.
func (h *Queue) PopFirst() (Task, error) {
	if h.IsEmpty() {
		return Task{}, fmt.Errorf("Queue is empty")
	}

	if (h.trimValue != -1) && (int(h.slice[0].task.Estimate) >= h.trimValue) {
		return Task{}, fmt.Errorf("Queue is empty")
	}

	val := heap.Pop(h)
	return val.(*heapRecord).task, nil
}

// String implements the Stringer interface
// Used mainly for testing
func (h *Queue) String() string {
	var b strings.Builder

	duplicate := NewHeapQueue()
	duplicate.slice = append(duplicate.slice, h.slice...)

	for val, err := duplicate.PopFirst(); err == nil; val, err = duplicate.PopFirst() {
		fmt.Fprintf(&b, " %d", val.Distance)
	}

	return fmt.Sprintf("tasks.Heap:%s", b.String())
}
