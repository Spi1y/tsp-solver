package iterator

import (
	"testing"

	"github.com/Spi1y/tsp-solver/solver2/types"
	"github.com/stretchr/testify/assert"
)

func TestIterator_NodesToVisit(t *testing.T) {
	tests := []struct {
		name    string
		size    types.Index
		path    []types.Index
		want    []types.Index
		wantErr bool
	}{
		{
			"err empty matrix",
			0, []types.Index{},
			[]types.Index{}, true,
		},
		{
			"err empty path",
			3, []types.Index{},
			[]types.Index{}, true,
		},
		{
			"err path not started from 0",
			3, []types.Index{2, 0},
			[]types.Index{}, true,
		},
		{
			"err path with wrong index",
			3, []types.Index{0, 8},
			[]types.Index{}, true,
		},
		{
			"err path too long",
			3, []types.Index{0, 3, 2, 1, 3},
			[]types.Index{}, true,
		},
		{
			"size 1 path 1",
			1, []types.Index{0},
			[]types.Index{}, false,
		},
		{
			"size 4 path 1",
			4, []types.Index{0},
			[]types.Index{1, 2, 3}, false,
		},
		{
			"size 4 path 2",
			4, []types.Index{0, 2},
			[]types.Index{1, 3}, false,
		},
		{
			"size 4 path 3",
			4, []types.Index{0, 2, 3},
			[]types.Index{1}, false,
		},
		{
			"size 4 path 4",
			4, []types.Index{0, 2, 3, 1},
			[]types.Index{}, false,
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
		size    types.Index
		path    []types.Index
		node    types.Index
		want    []types.Index
		wantErr bool
	}{
		{
			"err next node is visited",
			4, []types.Index{0, 2}, 2,
			nil, true,
		},
		{
			"err next node is out of bound",
			4, []types.Index{0, 2}, 5,
			nil, true,
		},
		{
			"err last node is not 0",
			4, []types.Index{0, 2, 1, 3}, 1,
			nil, true,
		},
		{
			"size 4 path 1",
			4, []types.Index{0}, 2,
			[]types.Index{0, 1, 3}, false,
		},
		{
			"size 4 path 2",
			4, []types.Index{0, 2}, 1,
			[]types.Index{0, 3}, false,
		},
		{
			"size 4 path 3",
			4, []types.Index{0, 2, 1}, 3,
			[]types.Index{0}, false,
		},
		{
			"size 4 path 4",
			4, []types.Index{0, 2, 1, 3}, 0,
			[]types.Index{}, false,
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
