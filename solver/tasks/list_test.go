package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_rawInsert111(t *testing.T) {
	list := &List{}
	task := &Task{}
	got := ""
	expected := ""

	t.Run("Insert to empty list", func(t *testing.T) {
		list.rawInsert(task, 1)
		got = list.String()
		expected = "tasks.List: 1"
		assert.Equal(t, got, expected, "After rawInsert list became %v, should be %v", got, expected)
	})

	list.rawInsert(task, 10)
	got = list.String()
	expected = "tasks.List: 1 10"
	assert.Equal(t, got, expected, "After rawInsert list became %v, should be %v", got, expected)

	list.rawInsert(task, 1)
	got = list.String()
	expected = "tasks.List: 1"
	assert.Equal(t, got, expected, "After 1st rawInsert list became %v, should be %v", got, expected)

	list.rawInsert(task, 1)
	got = list.String()
	expected = "tasks.List: 1"
	assert.Equal(t, got, expected, "After 1st rawInsert list became %v, should be %v", got, expected)
}

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
			"Insert to the front",
			10,
			"tasks.List: 10 5 1",
		},
		{
			"Insert into the middle",
			7,
			"tasks.List: 10 7 5 1",
		},
		{
			"Insert duplicate",
			5,
			"tasks.List: 10 7 5 5 1",
		},
	}

	list := &List{}
	task := &Task{}
	got := ""

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.rawInsert(task, tt.potential)
			got = list.String()
			assert.Equal(t, got, tt.expected)
		})
	}
}
