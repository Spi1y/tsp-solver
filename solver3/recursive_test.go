package solver3

import (
	"testing"

	"github.com/Spi1y/tsp-solver/solver2/iterator"
	"github.com/Spi1y/tsp-solver/solver2/tasks"
	"github.com/Spi1y/tsp-solver/solver2/types"
	"github.com/stretchr/testify/assert"
)

func TestSolver_SolveRecursively(t *testing.T) {
	tests := solverTestCases()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{
				matrix:               tt.distanceMatrix,
				bestSolution:         []types.Index{},
				bestSolutionDistance: 0,
				taskQueue:            tasks.NewHeapQueue(),
			}
			it := &iterator.Iterator{}
			it.Init(types.Index(len(tt.distanceMatrix)))

			it.SetPath([]types.Index{0})
			nextNodes := it.NodesToVisit()

			path, dist := s.solveRecursively(0, nextNodes)
			fullpath := make([]types.Index, 0)
			fullpath = append(fullpath, 0)
			fullpath = append(fullpath, path...)

			assert.Equal(t, tt.path, fullpath)
			assert.Equal(t, tt.dist, dist)
		})
	}
}
