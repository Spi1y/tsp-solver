package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap_Insert(t *testing.T) {
	tests := []struct {
		name      string
		distances []int
		expected  string
	}{
		{
			"Into empty list", []int{5},
			"tasks.Heap: 5",
		},
		{
			"Empty insert", []int{},
			"tasks.Heap: 5",
		},
		{
			"To the end", []int{1},
			"tasks.Heap: 1 5",
		},
		{
			"To the front", []int{10},
			"tasks.Heap: 1 5 10",
		},
		{
			"Into the middle", []int{7},
			"tasks.Heap: 1 5 7 10",
		},
		{
			"Duplicate", []int{5},
			"tasks.Heap: 1 5 5 7 10",
		},
		{
			"Bulk sorted", []int{12, 12, 6, 5, 5, 5, 4, 1, 0, 0},
			"tasks.Heap: 0 0 1 1 4 5 5 5 5 5 6 7 10 12 12",
		},
		{
			"Bulk mixed", []int{3, 8, 50, 50, 15, 0, 7, 6, 7, 0},
			"tasks.Heap: 0 0 0 0 1 1 3 4 5 5 5 5 5 6 6 7 7 7 8 10 12 12 15 50 50",
		},
	}
	list := NewHeapQueue()
	args := make([]*Task, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args = args[:0]
			for _, potential := range tt.distances {
				args = append(args, &Task{Distance: potential})
			}

			list.Insert(args)
			assert.Equal(t, tt.expected, list.String())
		})
	}
}
