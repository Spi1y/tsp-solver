package solver3

import (
	"errors"

	"github.com/Spi1y/tsp-solver/solver2/tasks"
	"github.com/Spi1y/tsp-solver/solver2/types"
)

// Solver is a TSP solver object. It is used to set a distance matrix and start
// calculations
type Solver struct {
	RecursiveThreshold types.Index

	// Distance matrix
	matrix [][]types.Distance
	// Tasks queue
	taskQueue *tasks.Queue

	// Current best solution
	bestSolution         []types.Index
	bestSolutionDistance types.Distance
}

// Solve solves the TSP problem with a given distance matrix.
func (s *Solver) Solve(m [][]types.Distance) ([]types.Index, types.Distance, error) {
	size := len(m)

	if size == 0 {
		return nil, 0, errors.New("Distance matrix is empty")
	}

	for i := range m {
		if len(m[i]) != size {
			return nil, 0, errors.New("Distance matrix is not square")
		}
	}

	s.matrix = m
	s.bestSolution = []types.Index{}
	s.bestSolutionDistance = 0
	s.taskQueue = tasks.NewHeapQueue()

	rootTask := tasks.Task{
		Path:     []types.Index{0},
		Distance: 0,
		Estimate: 0,
	}
	s.taskQueue.Insert([]tasks.Task{rootTask})

	s.solveParallel()

	return s.bestSolution, s.bestSolutionDistance, nil
}

func (s *Solver) newSolutionFound(path []types.Index, distance types.Distance) {
	if (s.bestSolutionDistance != 0) && (distance >= s.bestSolutionDistance) {
		return
	}

	s.bestSolution = path
	s.bestSolutionDistance = distance

	s.taskQueue.TrimTail(distance)
}
