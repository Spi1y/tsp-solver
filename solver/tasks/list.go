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
	Potential int
}

// List is a double-linked priority list of tasks
type List struct {
	First *ListRecord
	Last  *ListRecord

	// CurrentBestSolution is a current best found solution (if any)
	CurrentBestSolution []int
	// Distance of the CurrentBestSolution.
	// Used to delete obsolete solution paths from the end of the list
	CurrentBestDistance int

	// insertionQueue used for the optimization of bulk tasks insertion
	// See List.Insert(tasks)
	insertionQueue *List
}

// Insert inserts new tasks into the list
// Their potential is calculated and used to determine the position
// of the inserts
func (l *List) Insert(tasks ...*Task) {
	if l.insertionQueue == nil {
		l.insertionQueue = &List{}
	}

	// var potential int
	// for _, task := range tasks {
	// 	potential = task.ActualDistance + task.ProjectedDistance

	// }
}

// SolutionFound notifies the list of the finding of the new solution
// The list then checks if it`s better than the current one. If necessary,
// it will update it and cut obsolete solutions accordingly.
func (l *List) SolutionFound(solution []int, distance int) {

}

// Clear clears the list of all entries
func (l *List) Clear() {

}

// String implements the Stringer interface
// Used mainly for testing
func (l *List) String() string {
	var b strings.Builder
	for record := l.First; record != nil; record = record.next {
		fmt.Fprintf(&b, " %d", record.Potential)
	}
	return fmt.Sprintf("tasks.List:%s", b.String())
}

// rawInsert inserts one record into the list without optimizations,
// using simple interation. It is used to populate insertionQueue for
// optimized bulk insertion into the main list
func (l *List) rawInsert(task *Task, potential int) {
	record := &ListRecord{
		prev:      nil,
		next:      nil,
		Task:      task,
		Potential: potential,
	}

	if l.First == nil {
		l.First = record
		l.Last = record

		return
	}

	if l.First.Potential <= record.Potential {
		record.next = l.First
		l.First.prev = record
		l.First = record

		return
	}

	for checked := l.First; checked != nil; checked = checked.next {
		if checked.Potential <= record.Potential {
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
}
