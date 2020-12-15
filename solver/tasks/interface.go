package tasks

import "fmt"

// Queue is an interface to represent different collections of tasks: heap, linked list, etc
type Queue interface {
	fmt.Stringer

	Insert(tasks []*Task)
	PopFirst() *Task

	TrimTail(distance int)
	IsEmpty() bool
}

// QueueType is a container type for a different types of queues
type QueueType int

const (
	// QueueLinkedList is linked list queue
	QueueLinkedList QueueType = iota
	// QueueHeap is heap queue
	QueueHeap
)

// CreateQueue sreates and returns a queue of a requested type
func CreateQueue(t QueueType) Queue {
	switch t {
	case QueueLinkedList:
		return NewListQueue()
	case QueueHeap:
		return NewHeapQueue()
	}

	return nil
}
