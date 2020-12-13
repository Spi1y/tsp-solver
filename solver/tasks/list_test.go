package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_rawInsert(t *testing.T) {
	list := &List{}
	task := &Task{}

	tests := []struct {
		name      string
		potential int
		expected  string
	}{
		{
			"Into empty list", 5,
			"tasks.List: 5",
		},
		{
			"To the end", 1,
			"tasks.List: 1 5",
		},
		{
			"Duplicate end", 1,
			"tasks.List: 1 1 5",
		},
		{
			"To the front", 10,
			"tasks.List: 1 1 5 10",
		},
		{
			"Duplicate front", 10,
			"tasks.List: 1 1 5 10 10",
		},
		{
			"Into the middle", 7,
			"tasks.List: 1 1 5 7 10 10",
		},
		{
			"Duplicate middle", 5,
			"tasks.List: 1 1 5 5 7 10 10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.rawInsert(task, tt.potential)
			assert.Equal(t, tt.expected, list.String())
		})
	}
}

func TestList_Insert(t *testing.T) {
	tests := []struct {
		name       string
		potentials []int
		expected   string
	}{
		{
			"Into empty list", []int{5},
			"tasks.List: 5",
		},
		{
			"Empty insert", []int{},
			"tasks.List: 5",
		},
		{
			"To the end", []int{1},
			"tasks.List: 1 5",
		},
		{
			"To the front", []int{10},
			"tasks.List: 1 5 10",
		},
		{
			"Into the middle", []int{7},
			"tasks.List: 1 5 7 10",
		},
		{
			"Duplicate", []int{5},
			"tasks.List: 1 5 5 7 10",
		},
		{
			"Bulk sorted", []int{12, 12, 6, 5, 5, 5, 4, 1, 0, 0},
			"tasks.List: 0 0 1 1 4 5 5 5 5 5 6 7 10 12 12",
		},
		{
			"Bulk mixed", []int{3, 8, 50, 50, 15, 0, 7, 6, 7, 0},
			"tasks.List: 0 0 0 0 1 1 3 4 5 5 5 5 5 6 6 7 7 7 8 10 12 12 15 50 50",
		},
	}
	list := &List{}
	args := make([]*Task, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args = args[:0]
			for _, potential := range tt.potentials {
				args = append(args, &Task{Distance: potential})
			}

			list.BulkPush(args)
			assert.Equal(t, tt.expected, list.String())
		})
	}
}

func TestList_TrimTail(t *testing.T) {
	list := &List{}
	initlist := []int{3, 4, 5, 5, 5, 5, 5, 6, 6, 7, 7, 7, 8, 10, 12, 12, 15, 50, 50}
	args := make([]*Task, len(initlist))
	for i, potential := range initlist {
		args[i] = &Task{Distance: potential}
	}
	list.BulkPush(args)

	tests := []struct {
		name      string
		potential int
		expected  string
	}{
		{
			"No trim", 100,
			"tasks.List: 3 4 5 5 5 5 5 6 6 7 7 7 8 10 12 12 15 50 50",
		},
		{
			"Simple", 15,
			"tasks.List: 3 4 5 5 5 5 5 6 6 7 7 7 8 10 12 12",
		},
		{
			"Simple-2", 5,
			"tasks.List: 3 4",
		},
		{
			"Full", 2,
			"tasks.List:",
		},
		{
			"On empty list", 1,
			"tasks.List:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.TrimTail(tt.potential)
			assert.Equal(t, tt.expected, list.String())
		})
	}
}

func TestList_IsEmpty(t *testing.T) {
	emptyList := &List{}
	notEmptyList := &List{}
	notEmptyList.rawInsert(&Task{}, 0)

	tests := []struct {
		name string
		l    *List
		want bool
	}{
		{"Empty", emptyList, true},
		{"Not empty", notEmptyList, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsEmpty(); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestList_GetFirst(t *testing.T) {
	list := &List{}
	tasks := []*Task{
		{Distance: 1},
		{Distance: 5},
		{Distance: 10},
	}
	list.BulkPush(tasks)

	tests := []struct {
		name      string
		want      *Task
		wantEmpty bool
	}{
		{"[0]", tasks[0], false},
		{"[1]", tasks[1], false},
		{"[2]", tasks[2], true},
		{"nil", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list.Pop()
			assert.Same(t, tt.want, got)
			assert.Equal(t, tt.wantEmpty, list.IsEmpty())
		})
	}
}
