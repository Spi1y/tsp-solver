package tasks

import (
	"fmt"
	"strings"
)

// ListRecord is a record in double-linked priority list of tasks
type ListRecord struct {
	prev *ListRecord
	next *ListRecord

	Task *Task
	// Estimation of Task potential to lead to the best solution
	// This is a sum of Task.ActualDistance and Task.ProjectedDistance
	Distance int
}

// List is a double-linked sorted list of tasks
type List struct {
	First *ListRecord
	Last  *ListRecord

	// insertionQueue used for the optimization of bulk tasks insertion
	// See List.Insert(tasks)
	insertionQueue *List
}

// NewListQueue creates and returns a new list
func NewListQueue() *List {
	return &List{}
}

// Insert inserts new tasks into the list
// Their potential is calculated and used to determine the position
// of the inserts
func (l *List) Insert(tasks []*Task) {
	// A quick path for an empty insertion
	if len(tasks) == 0 {
		return
	}

	// Initialize insertion queue if necessary
	if l.insertionQueue == nil {
		l.insertionQueue = &List{}
	}

	// Populate insertion queue and defer clearing
	for _, task := range tasks {
		l.insertionQueue.rawInsert(task, task.Distance)
	}
	defer l.insertionQueue.clear()

	// A quick path for an empty list
	if l.First == nil {
		l.First = l.insertionQueue.First
		l.Last = l.insertionQueue.Last

		return
	}

	insertion := l.insertionQueue.First

	// Check if new inserts have to go in the head of the list
	for (insertion != nil) && (l.First.Distance >= insertion.Distance) {
		// Extract the record from the insertion queue
		newRecord := insertion
		insertion = insertion.next

		// Install it as a new head
		newRecord.next = l.First
		l.First.prev = newRecord
		l.First = newRecord
	}

	// Iterate over list records, inserting as needed
	for listRecord := l.First; (listRecord != nil) && (insertion != nil); listRecord = listRecord.next {
		// Iterate over remaining insertion queue, if there is suitable insertions to go before the listRecord
		for (insertion != nil) && (listRecord.Distance >= insertion.Distance) {
			// Extract the record from the insertion queue
			newRecord := insertion
			insertion = insertion.next

			// Install it before the current list record
			listRecord.prev.next = newRecord
			newRecord.prev = listRecord.prev

			newRecord.next = listRecord
			listRecord.prev = newRecord
		}
	}

	// If some insertion records remain, they should go to the end
	if insertion != nil {
		insertion.prev = l.Last
		l.Last.next = insertion

		l.Last = l.insertionQueue.Last
	}
}

// TrimTail trims records from the tail of the list with distance greater
// than the given argument
func (l *List) TrimTail(distance int) {
	// A quick path for an empty list
	if l.First == nil {
		return
	}

	// A quick path for a tail not suitable for trimming
	if l.Last.Distance < distance {
		return
	}

	// A quick path for a full list trim
	if l.First.Distance >= distance {
		l.clear()
		return
	}

	for checked := l.Last; checked != nil; checked = checked.prev {
		if checked.Distance < distance {
			l.Last = checked
			l.Last.next = nil
			break
		}
	}
}

// IsEmpty checks if there is no records in the list.
func (l *List) IsEmpty() bool {
	if l.First == nil {
		return true
	}

	return false
}

// PopFirst gets the task from the first record in the list and
// removes it from the list.
// If list is empty, it returns nil.
func (l *List) PopFirst() *Task {
	if l.First == nil {
		return nil
	}

	record := l.First
	l.First = l.First.next

	return record.Task
}

// String implements the Stringer interface
// Used mainly for testing
func (l *List) String() string {
	var b strings.Builder
	for record := l.First; record != nil; record = record.next {
		fmt.Fprintf(&b, " %d", record.Distance)
	}
	return fmt.Sprintf("tasks.List:%s", b.String())
}

// rawInsert inserts one record into the list without optimizations,
// using simple interation. It is used to populate insertionQueue for
// optimized bulk insertion into the main list
func (l *List) rawInsert(task *Task, potential int) {
	record := &ListRecord{
		prev:     nil,
		next:     nil,
		Task:     task,
		Distance: potential,
	}

	// A quick path for an empty list
	if l.First == nil {
		l.First = record
		l.Last = record

		return
	}

	// Check if new insert have to go in the head of the list
	if l.First.Distance >= record.Distance {
		record.next = l.First
		l.First.prev = record
		l.First = record

		return
	}

	// Iterate over list records, seeking the suitable insert position
	for checked := l.First; checked != nil; checked = checked.next {
		if checked.Distance >= record.Distance {
			checked.prev.next = record
			record.prev = checked.prev

			record.next = checked
			checked.prev = record

			return
		}
	}

	// We have not found an insertion position, which means we have the smallest
	// available potential and have to insert in the end
	record.prev = l.Last
	l.Last.next = record
	l.Last = record
}

// clear clears the list of all records
func (l *List) clear() {
	l.First = nil
	l.Last = nil
}
