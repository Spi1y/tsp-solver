package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_rawInsert(t *testing.T) {
	tests := []struct {
		name      string
		potential int
		expected  string
	}{
		{
			"Insert into empty list",
			5,
			"tasks.List: 5",
		},
		{
			"Insert to the end",
			1,
			"tasks.List: 5 1",
		},
		{
			"Duplicate end",
			1,
			"tasks.List: 5 1 1",
		},
		{
			"Insert to the front",
			10,
			"tasks.List: 10 5 1 1",
		},
		{
			"Duplicate front",
			10,
			"tasks.List: 10 10 5 1 1",
		},
		{
			"Insert into the middle",
			7,
			"tasks.List: 10 10 7 5 1 1",
		},
		{
			"Duplicate middle",
			5,
			"tasks.List: 10 10 7 5 5 1 1",
		},
		{
			"Second duplicate middle",
			5,
			"tasks.List: 10 10 7 5 5 5 1 1",
		},
	}

	list := &List{}
	task := &Task{}

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
			"Insert into empty list",
			[]int{5},
			"tasks.List: 5",
		},
		{
			"Empty insert",
			[]int{},
			"tasks.List: 5",
		},
		{
			"Insert to the end",
			[]int{1},
			"tasks.List: 5 1",
		},
		{
			"Insert to the front",
			[]int{10},
			"tasks.List: 10 5 1",
		},
		{
			"Insert into the middle",
			[]int{7},
			"tasks.List: 10 7 5 1",
		},
		{
			"Insert duplicate",
			[]int{5},
			"tasks.List: 10 7 5 5 1",
		},
		{
			"Bulk insert sorted",
			[]int{12, 6, 5, 5, 5, 4, 1, 0, 0},
			"tasks.List: 12 10 7 6 5 5 5 5 5 4 1 1 0 0",
		},
		{
			"Bulk insert mixed",
			[]int{3, 8, 50, 15, 7, 6, 7, 0},
			"tasks.List: 50 15 12 10 8 7 7 7 6 6 5 5 5 5 5 4 3 1 1 0 0 0",
		},
	}
	list := &List{}
	args := make([]*Task, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args = args[:0]
			for _, potential := range tt.potentials {
				args = append(args, &Task{ActualDistance: potential})
			}

			list.Insert(args)
			assert.Equal(t, tt.expected, list.String())
		})
	}
}

func TestList_TrimTail(t *testing.T) {
	tests := []struct {
		name      string
		potential int
		expected  string
	}{
		{
			"Empty trim",
			-1,
			"tasks.List: 50 15 12 10 8 7 7 7 6 6 5 5 5 5 5 4 3 1 1 0 0 0",
		},
		{
			"Simple trim",
			0,
			"tasks.List: 50 15 12 10 8 7 7 7 6 6 5 5 5 5 5 4 3 1 1",
		},
		{
			"Simple trim - 2",
			5,
			"tasks.List: 50 15 12 10 8 7 7 7 6 6",
		},
		{
			"Full trim",
			100,
			"tasks.List:",
		},
		{
			"Trim on empty list",
			100,
			"tasks.List:",
		},
	}

	list := &List{}
	initlist := []int{50, 15, 12, 10, 8, 7, 7, 7, 6, 6, 5, 5, 5, 5, 5, 4, 3, 1, 1, 0, 0, 0}
	args := make([]*Task, 0)
	for _, potential := range initlist {
		args = append(args, &Task{ActualDistance: potential})
	}
	list.Insert(args)

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
		{
			"Empty",
			emptyList,
			true,
		},
		{
			"Not empty",
			notEmptyList,
			false,
		},
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
		{ActualDistance: 10},
		{ActualDistance: 5},
		{ActualDistance: 1},
	}
	list.Insert(tasks)

	tests := []struct {
		name      string
		want      *Task
		wantEmpty bool
	}{
		{
			"Get 10",
			tasks[0],
			false,
		},
		{
			"Get 5",
			tasks[1],
			false,
		},
		{
			"Get 1",
			tasks[2],
			true,
		},
		{
			"Get nil",
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list.GetFirst()
			assert.Same(t, tt.want, got)
			assert.Equal(t, tt.wantEmpty, list.IsEmpty())
		})
	}
}
