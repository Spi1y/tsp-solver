package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_NodesToVisit(t *testing.T) {
	tests := []struct {
		name    string
		size    uint8
		path    []uint8
		want    []uint8
		wantErr bool
	}{
		{
			"err empty matrix",
			0, []uint8{},
			[]uint8{}, true,
		},
		{
			"err empty path",
			3, []uint8{},
			[]uint8{}, true,
		},
		{
			"err path not started from 0",
			3, []uint8{2, 0},
			[]uint8{}, true,
		},
		{
			"err path with wrong index",
			3, []uint8{0, 8},
			[]uint8{}, true,
		},
		{
			"err path too long",
			3, []uint8{0, 3, 2, 1, 3},
			[]uint8{}, true,
		},
		{
			"size 1 path 1",
			1, []uint8{0},
			[]uint8{}, false,
		},
		{
			"size 4 path 1",
			4, []uint8{0},
			[]uint8{1, 2, 3}, false,
		},
		{
			"size 4 path 2",
			4, []uint8{0, 2},
			[]uint8{1, 3}, false,
		},
		{
			"size 4 path 3",
			4, []uint8{0, 2, 3},
			[]uint8{1}, false,
		},
		{
			"size 4 path 4",
			4, []uint8{0, 2, 3, 1},
			[]uint8{}, false,
		},
	}

	i := &Iterator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i.Init(tt.size)
			err := i.SetPath(tt.path)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err == nil {
				got := i.NodesToVisit()
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIterator_ColumnsToIterate(t *testing.T) {
	tests := []struct {
		name    string
		size    uint8
		path    []uint8
		node    uint8
		want    []uint8
		wantErr bool
	}{
		{
			"err next node is visited",
			4, []uint8{0, 2}, 2,
			nil, true,
		},
		{
			"err next node is out of bound",
			4, []uint8{0, 2}, 5,
			nil, true,
		},
		{
			"err last node is not 0",
			4, []uint8{0, 2, 1, 3}, 1,
			nil, true,
		},
		{
			"size 4 path 1",
			4, []uint8{0}, 2,
			[]uint8{0, 1, 3}, false,
		},
		{
			"size 4 path 2",
			4, []uint8{0, 2}, 1,
			[]uint8{0, 3}, false,
		},
		{
			"size 4 path 3",
			4, []uint8{0, 2, 1}, 3,
			[]uint8{0}, false,
		},
		{
			"size 4 path 4",
			4, []uint8{0, 2, 1, 3}, 0,
			[]uint8{}, false,
		},
	}

	i := &Iterator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i.Init(tt.size)
			err := i.SetPath(tt.path)
			assert.NoError(t, err)

			if err == nil {
				got, err := i.ColsToIterate(tt.node)
				assert.Equal(t, tt.want, got)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
