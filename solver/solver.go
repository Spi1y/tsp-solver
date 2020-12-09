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
}

// Solve solves the TSP problem with a given distance matrix.
func (s *Solver) Solve() ([]int, error) {
	size := len(s.DistanceMatrix)

	if size == 0 {
		return nil, errors.New("empty matrix")
	}

	// Mark 0-0 path as processed to correctly detect it in the solveTask
	// It looks ugly, but it is necessary due to the specifications of the
	// task - circular path with start and end in 0. So 0 node is special and
	// we treat it as such
	for i := range s.DistanceMatrix {
		s.DistanceMatrix[i][i] = -1
	}

	//rotNode := solutionNode{
	// _ = solutionNode{
	// 	parent:        nil,
	// 	point:         0,
	// 	projectedCost: 0,
	// 	actualCost:    0,
	// 	matrix:        nil,
	// }

	return nil, nil
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
		if m[0][nextNode] == -1 {
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
			s.bestSolution = newPath
			s.bestSolutionDistance = fullDistance

			// There won`t be new tasks
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
