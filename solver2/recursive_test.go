package solver2

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
				buffer:               make([]types.Distance, len(tt.distanceMatrix)),
				taskQueue:            tasks.NewHeapQueue(),
				iterator:             &iterator.Iterator{},
			}
			s.iterator.Init(types.Index(len(tt.distanceMatrix)))

			s.iterator.SetPath([]types.Index{0})
			nextNodes := s.iterator.NodesToVisit()

			path, dist := s.SolveRecursively(0, nextNodes)
			fullpath := make([]types.Index, 0)
			fullpath = append(fullpath, 0)
			fullpath = append(fullpath, path...)

			assert.Equal(t, tt.path, fullpath)
			assert.Equal(t, tt.dist, dist)
		})
	}
}
