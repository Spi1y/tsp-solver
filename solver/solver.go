package solver

import (
	"errors"

	"github.com/Spi1y/tsp-solver/solver/matrix"
	"github.com/Spi1y/tsp-solver/solver/tasks"
)

// Solver is an object used to encapulate internal state of
// the algorithm
type Solver struct {
	DistanceMatrix matrix.Matrix

	bestSolution         []int
	bestSolutionDistance int

	tasks *tasks.List
}

// Solve solves the TSP problem with a given distance matrix.
func (s *Solver) Solve() ([]int, int, error) {
	size := len(s.DistanceMatrix)

	if size == 0 {
		return nil, 0, errors.New("empty matrix")
	}

	s.bestSolution = []int{}
	s.bestSolutionDistance = 0

	rootMatrix := s.DistanceMatrix.Copy()
	basePathCost := rootMatrix.Normalize()
	// Mark 0-0 path as processed to correctly skip it in firther calculations
	rootMatrix[0][0] = -1
	rootTask := &tasks.Task{
		Path:           []int{0},
		Distance:       basePathCost,
		DistanceMatrix: rootMatrix,
	}
	s.tasks = &tasks.List{}

	newTasks := s.solveTask(rootTask)
	s.tasks.BulkPush(newTasks)

	for !s.tasks.IsEmpty() {
		task := s.tasks.Pop()
		newTasks := s.solveTask(task)
		s.tasks.BulkPush(newTasks)
	}

	return s.bestSolution, s.bestSolutionDistance, nil
}

func (s *Solver) newSolutionFound(path []int, distance int) {
	if (s.bestSolutionDistance != 0) && (distance >= s.bestSolutionDistance) {
		return
	}

	s.bestSolution = path
	s.bestSolutionDistance = distance

	s.tasks.TrimTail(distance)
}

func (s *Solver) solveTask(task *tasks.Task) []*tasks.Task {
	if task == nil {
		return nil
	}

	m := task.DistanceMatrix
	// The len(m) takes into account the 0 node, which should be
	// processed only in special case of closingNode. Thus, minus one.
	nodesTotal := len(m) - 1
	// The task.Path always includes the first 0 node, which is not
	// accounted for in the nodesTotal. Thus, minus one.
	nodesTraversed := len(task.Path) - 1
	currentNode := task.Path[nodesTraversed]
	// Check if this is the last node of the path
	closingNode := (nodesTraversed == nodesTotal-1)

	newTasks := make([]*tasks.Task, 0, nodesTotal-nodesTraversed)
	for nextNode := range m {
		// Skip already visited nodes
		if m[nextNode][0] == -1 {
			continue
		}

		nodeMatrix := task.DistanceMatrix.Copy()
		nodeMatrix.CutNode(currentNode, nextNode, closingNode)
		normalizationCost := nodeMatrix.Normalize()

		distanceToNode := m[currentNode][nextNode]
		fullDistance := task.Distance + distanceToNode + normalizationCost

		if (len(s.bestSolution) != 0) && (fullDistance >= s.bestSolutionDistance) {
			// And here is where we get to cutting from "branch and cut."
			continue
		}

		newPath := make([]int, 0, nodesTraversed+1)
		newPath = append(newPath, task.Path...)
		newPath = append(newPath, nextNode)

		if closingNode {
			// The path is finished and it`s better than the current best one
			// Update the solver state
			s.newSolutionFound(newPath, fullDistance)
			return nil
		}

		newTasks = append(newTasks, &tasks.Task{
			Path:           newPath,
			Distance:       fullDistance,
			DistanceMatrix: nodeMatrix,
		})
	}

	return newTasks
}
