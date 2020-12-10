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
		path           []int
		dist           int
		wantErr        bool
	}{
		{
			"nil",
			matrix.ConvertToMatrix([][]int{}),
			nil,
			0,
			true,
		},
		{
			"Normal - 2 points",
			matrix.ConvertToMatrix([][]int{
				{0, 1, 9},
				{9, 0, 1},
				{1, 9, 0},
			}),
			[]int{0, 1, 2},
			3,
			false,
		},
		{
			"Normal - 2 points",
			matrix.ConvertToMatrix([][]int{
				{0, 1, 9},
				{9, 0, 1},
				{1, 9, 0},
			}),
			[]int{0, 1, 2},
			3,
			false,
		},
		{
			"Real case - 4 points",
			matrix.ConvertToMatrix([][]int{
				{-1, 15_147, 4_596, 10_263, 5_482},
				{17_465, -1, 19_314, 21_477, 20_619},
				{4_643, 20_347, -1, 6_918, 1_340},
				{10_506, 21_310, 7_257, -1, 6_089},
				{6_585, 20_577, 1_340, 6_199, -1},
			}),
			[]int{0, 1, 3, 4, 2},
			48_696,
			false,
		},
		{
			"Real case - 7 points",
			matrix.ConvertToMatrix([][]int{
				{-1, 15147, 21742, 12730, 18594, 6147, 6955, 10000},
				{17465, -1, 30524, 22534, 27376, 20763, 15326, 21214},
				{23594, 43627, -1, 16165, 9604, 21957, 18560, 21180},
				{11103, 22595, 16255, -1, 10210, 5909, 7880, 3274},
				{19133, 27796, 9754, 10054, -1, 12856, 14099, 10486},
				{6155, 21069, 23218, 7694, 14520, -1, 5419, 4964},
				{5736, 14952, 18081, 8492, 14933, 6300, -1, 7172},
				{10801, 21605, 17131, 4504, 11197, 3615, 6890, -1},
			}),
			[]int{0, 1, 6, 2, 4, 3, 7, 5},
			81_256,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}
			s.DistanceMatrix = tt.distanceMatrix
			path, dist, err := s.Solve()

			assert.Equal(t, tt.path, path)
			assert.Equal(t, tt.dist, dist)
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
		bestDistOut  int
		bestPathOut  []int
	}{
		{
			"nil",
			0,
			nil,
			nil,
			0,
			nil,
		},
		{
			"First step task",
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
					DistanceMatrix: matrix.ConvertToMatrix([][]int{
						{-1, -1, -1, -1},
						{-1, -1, 0, 8},
						{8, -1, 8, 0},
						{0, -1, 8, 8},
					}),
				},
				{
					Path:     []int{0, 2},
					Distance: 20,
					DistanceMatrix: matrix.ConvertToMatrix([][]int{
						{-1, -1, -1, -1},
						{0, 0, -1, 0},
						{-1, 8, -1, 0},
						{0, 8, -1, 8},
					}),
				},
				{
					Path:     []int{0, 3},
					Distance: 28,
					DistanceMatrix: matrix.ConvertToMatrix([][]int{
						{-1, -1, -1, -1},
						{8, 8, 0, -1},
						{0, 0, 0, -1},
						{-1, 0, 0, -1},
					}),
				},
			},
			0,
			nil,
		},
		{
			"Second step task",
			0,
			&tasks.Task{
				Path:     []int{0, 1},
				Distance: 4,
				DistanceMatrix: matrix.ConvertToMatrix([][]int{
					{-1, -1, -1, -1},
					{-1, -1, 0, 8},
					{8, -1, 8, 0},
					{0, -1, 8, 8},
				}),
			},
			[]*tasks.Task{
				{
					Path:     []int{0, 1, 2},
					Distance: 4,
					DistanceMatrix: matrix.ConvertToMatrix([][]int{
						{-1, -1, -1, -1},
						{-1, -1, -1, -1},
						{-1, -1, -1, 0},
						{0, -1, -1, 8},
					}),
				},
				{
					Path:     []int{0, 1, 3},
					Distance: 28,
					DistanceMatrix: matrix.ConvertToMatrix([][]int{
						{-1, -1, -1, -1},
						{-1, -1, -1, -1},
						{0, -1, 0, -1},
						{-1, -1, 0, -1},
					}),
				},
			},
			0,
			nil,
		},
		{
			"Last step task",
			0,
			&tasks.Task{
				Path:     []int{0, 1, 2},
				Distance: 4,
				DistanceMatrix: matrix.ConvertToMatrix([][]int{
					{-1, -1, -1, -1},
					{-1, -1, -1, -1},
					{-1, -1, -1, 0},
					{0, -1, -1, 8},
				}),
			},
			nil,
			4,
			[]int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}
			tasks := s.solveTask(tt.task)

			assert.Equal(t, tt.want, tasks)
			assert.Equal(t, tt.bestDistOut, s.bestSolutionDistance)
			assert.Equal(t, tt.bestPathOut, s.bestSolution)
		})
	}
}
