package solver2

import (
	"errors"

	"github.com/Spi1y/tsp-solver/solver2/iterator"
	"github.com/Spi1y/tsp-solver/solver2/tasks"
	"github.com/Spi1y/tsp-solver/solver2/types"
)

// Solver is a TSP solver object. It is used to set a distance matrix and start
// calculations
type Solver struct {
	RecursiveThreshold types.Index

	// Distance matrix
	matrix [][]types.Distance
	// Iterator (see package docs)
	iterator *iterator.Iterator
	// Tasks queue
	taskQueue *tasks.Queue
	// Temporary buffer to optimize normalization
	buffer []types.Distance

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
	s.buffer = make([]types.Distance, size)
	s.taskQueue = tasks.NewHeapQueue()
	s.iterator = &iterator.Iterator{}
	s.iterator.Init(types.Index(size))

	newTasks := make([]tasks.Task, size)

	rootTask := tasks.Task{
		Path:     []types.Index{0},
		Distance: 0,
		Estimate: 0,
	}
	s.taskQueue.Insert([]tasks.Task{rootTask})

	for task, err := s.taskQueue.PopFirst(); err == nil; task, err = s.taskQueue.PopFirst() {
		count, err := s.solveTask(task, newTasks)
		if err != nil {
			return nil, 0, err
		}

		s.taskQueue.Insert(newTasks[:count])
	}

	return s.bestSolution, s.bestSolutionDistance, nil
}

func (s *Solver) solveTask(t tasks.Task, newTasks []tasks.Task) (int, error) {
	// TODO - try aggressive approach with full path first

	err := s.iterator.SetPath(t.Path)
	if err != nil {
		return 0, err
	}
	nextNodes := s.iterator.NodesToVisit()
	rows := s.iterator.RowsToIterate()

	currNode := t.Path[len(t.Path)-1]
	nodesLeft := len(nextNodes)

	if nodesLeft <= int(s.RecursiveThreshold) {
		tailpath, taildistance := s.solveRecursively(currNode, nextNodes)
		path := make([]types.Index, len(t.Path), len(t.Path)+len(tailpath))
		copy(path, t.Path)
		path = append(path, tailpath...)
		distance := t.Distance + taildistance

		s.newSolutionFound(path, distance)
		newTasks = newTasks[:0]
		return 0, nil
	}

	if nodesLeft == 1 {
		// Final node, calculating return distance to root node
		// and notifying solver about found solution
		finalNode := nextNodes[0]

		path := make([]types.Index, len(t.Path), len(t.Path)+2)
		copy(path, t.Path)
		path = append(path, finalNode, 0)
		distance := t.Distance + s.matrix[currNode][finalNode] + s.matrix[finalNode][0]
		s.newSolutionFound(path, distance)
		newTasks = newTasks[:0]
		return 0, nil
	}

	newPathLen := len(t.Path) + 1
	pathsSlice := make([]types.Index, nodesLeft*newPathLen)

	for i, nextNode := range nextNodes {

		var estimate types.Distance
		cols, err := s.iterator.ColsToIterate(nextNode)
		if err != nil {
			return 0, err
		}

		for rowIndex, row := range rows {
			rowSlice := s.matrix[row]

			min := rowSlice[0]
			var val types.Distance

			// First pass to calculate row minimum
			for _, col := range cols {
				if row == col {
					continue
				}

				val := rowSlice[col]
				if min > val {
					min = val
				}
			}
			estimate += min

			// Second pass to update column minimums in the buffer
			if rowIndex == 0 {
				// Fast path for a first row
				for colIndex, col := range cols {
					if row == col {
						continue
					}
					// First row, minimum values are set without comparison
					s.buffer[colIndex] = rowSlice[col] - min
				}
				continue
			}

			// Normal path for other rows
			for colIndex, col := range cols {
				if row == col {
					continue
				}
				val = rowSlice[col] - min

				// Values are updated as needed
				if s.buffer[colIndex] > val {
					s.buffer[colIndex] = val
				}
			}
		}

		// Final pass on buffer to sum column minimums
		for colIndex := range cols {
			estimate += s.buffer[colIndex]
		}

		path := pathsSlice[i*newPathLen : (i+1)*newPathLen]
		copy(path, t.Path)
		path[newPathLen-1] = nextNode

		distance := t.Distance + s.matrix[currNode][nextNode]

		newTasks[i].Path = path
		newTasks[i].Distance = distance
		newTasks[i].Estimate = distance + estimate
	}

	return nodesLeft, nil
}

func (s *Solver) newSolutionFound(path []types.Index, distance types.Distance) {
	if (s.bestSolutionDistance != 0) && (distance >= s.bestSolutionDistance) {
		return
	}

	s.bestSolution = path
	s.bestSolutionDistance = distance

	s.taskQueue.TrimTail(distance)
}
