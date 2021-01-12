package solver2

import (
	"testing"

	"github.com/Spi1y/tsp-solver/solver2/types"
	"github.com/stretchr/testify/assert"
)

func TestSolverSolve(t *testing.T) {
	tests := solverTestCases()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}

			path, dist, err := s.Solve(tt.distanceMatrix)

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

func solverTestCases() []*solveTestCase {
	result := []*solveTestCase{}
	result = append(result, solveTestCase2Points())
	result = append(result, solveTestCase3Points())
	result = append(result, solveTestCase4Points())
	result = append(result, solveTestCase7Points())

	return result
}

type solveTestCase struct {
	name           string
	distanceMatrix [][]types.Distance
	path           []types.Index
	dist           types.Distance
	wantErr        bool
}

func solveTestCase2Points() *solveTestCase {
	return &solveTestCase{
		"Normal - 2 points",
		[][]types.Distance{
			{0, 1, 9},
			{9, 0, 1},
			{1, 9, 0},
		},
		[]types.Index{0, 1, 2, 0},
		3,
		false,
	}
}

func solveTestCase3Points() *solveTestCase {
	return &solveTestCase{
		"Normal - 3 points",
		[][]types.Distance{
			{0, 1, 9, 9},
			{9, 0, 9, 1},
			{1, 9, 0, 9},
			{9, 9, 1, 0},
		},
		[]types.Index{0, 1, 3, 2, 0},
		4,
		false,
	}
}

func solveTestCase4Points() *solveTestCase {
	return &solveTestCase{
		"Real case - 4 points",
		[][]types.Distance{
			{0, 15_147, 4_596, 10_263, 5_482},
			{17_465, 0, 19_314, 21_477, 20_619},
			{4_643, 20_347, 0, 6_918, 1_340},
			{10_506, 21_310, 7_257, 0, 6_089},
			{6_585, 20_577, 1_340, 6_199, 0},
		},
		[]types.Index{0, 1, 3, 4, 2, 0},
		48_696,
		false,
	}
}

func solveTestCase7Points() *solveTestCase {
	return &solveTestCase{
		"Real case - 7 points",
		[][]types.Distance{
			{0, 15147, 21742, 12730, 18594, 6147, 6955, 10000},
			{17465, 0, 30524, 22534, 27376, 20763, 15326, 21214},
			{23594, 43627, 0, 16165, 9604, 21957, 18560, 21180},
			{11103, 22595, 16255, 0, 10210, 5909, 7880, 3274},
			{19133, 27796, 9754, 10054, 0, 12856, 14099, 10486},
			{6155, 21069, 23218, 7694, 14520, 0, 5419, 4964},
			{5736, 14952, 18081, 8492, 14933, 6300, 0, 7172},
			{10801, 21605, 17131, 4504, 11197, 3615, 6890, 0},
		},
		[]types.Index{0, 1, 6, 2, 4, 3, 7, 5, 0},
		81_256,
		false,
	}
}
