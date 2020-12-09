package solver

import (
	"testing"

	"github.com/Spi1y/tsp-solver/solver/matrix"
	"github.com/Spi1y/tsp-solver/solver/tasks"
	"github.com/stretchr/testify/assert"
)

func TestSolver_Solve(t *testing.T) {
	tests := []struct {
		name           string
		distanceMatrix matrix.Matrix
		want           []int
		wantErr        bool
	}{
		{
			"nil",
			matrix.ConvertToMatrix([][]int{}),
			nil,
			true,
		},
		{
			"Normal - 2 points",
			matrix.ConvertToMatrix([][]int{
				{0, 1, 9},
				{9, 0, 1},
				{1, 9, 0},
			}),
			[]int{0, 1, 2, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}
			s.DistanceMatrix = tt.distanceMatrix
			got, err := s.Solve()

			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSolver_solveTask(t *testing.T) {
	tests := []struct {
		name         string
		bestDistance int
		task         *tasks.Task
		want         []*tasks.Task
	}{
		{
			"nil",
			0,
			nil,
			nil,
		},
		{
			"First task",
			0,
			&tasks.Task{
				Path:     []int{0},
				Distance: 4,
				DistanceMatrix: matrix.ConvertToMatrix([][]int{
					{-1, 0, 8, 8},
					{8, 8, 0, 8},
					{8, 8, 8, 0},
					{0, 8, 8, 8},
				}),
			},
			[]*tasks.Task{
				{
					Path:     []int{0, 1},
					Distance: 4,
					DistanceMatrix: [][]int{
						{-1, -1, -1, -1},
						{-1, -1, 0, 8},
						{8, -1, 8, 0},
						{0, -1, 8, 8},
					},
				},
				{
					Path:     []int{0, 2},
					Distance: 20,
					DistanceMatrix: [][]int{
						{-1, -1, -1, -1},
						{0, 0, -1, 0},
						{-1, 8, -1, 0},
						{0, 8, -1, 8},
					},
				},
				{
					Path:     []int{0, 3},
					Distance: 28,
					DistanceMatrix: [][]int{
						{-1, -1, -1, -1},
						{8, 8, 0, -1},
						{0, 0, 0, -1},
						{-1, 0, 0, -1},
					},
				},
			},
		},
		// TODO - Implement further tests
		// {
		// 	"Middle task",
		// 	0,
		// 	&tasks.Task{},
		// 	[]*tasks.Task{{}},
		// },
		// {
		// 	"Last node",
		// 	0,
		// 	&tasks.Task{},
		// 	[]*tasks.Task{{}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}
			tasks := s.solveTask(tt.task)

			assert.Equal(t, tt.want, tasks)
		})
	}
}
