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

	rootTask := &tasks.Task{
		Path:     []types.Index{0},
		Distance: 0,
		Estimate: 0,
	}
	s.taskQueue.Insert([]*tasks.Task{rootTask})

	for !s.taskQueue.IsEmpty() {
		task := s.taskQueue.PopFirst()
		newTasks, err := s.solveTask(task)
		if err != nil {
			return nil, 0, err
		}
		s.taskQueue.Insert(newTasks)
	}

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

func (s *Solver) solveTask(t *tasks.Task) ([]*tasks.Task, error) {
	err := s.iterator.SetPath(t.Path)
	if err != nil {
		return nil, err
	}
	nextNodes := s.iterator.NodesToVisit()
	rows := s.iterator.RowsToIterate()

	currNode := t.Path[len(t.Path)-1]
	nodesLeft := len(nextNodes)
	if nodesLeft == 0 {
		// Final node, calculating return distance to root node
		// and notifying solver about found solution
		path := append(t.Path, 0)
		distance := t.Distance + s.matrix[currNode][0]
		s.newSolutionFound(path, distance)
		return nil, nil
	}

	newTasks := make([]*tasks.Task, nodesLeft)

	for i, nextNode := range nextNodes {

		var estimate types.Distance
		cols, err := s.iterator.ColsToIterate(nextNode)
		if err != nil {
			return nil, err
		}

		for rowIndex, row := range rows {
			min := s.matrix[row][0]
			var val types.Distance

			// First pass to calculate row minimum
			for _, col := range cols {
				if row == col {
					continue
				}

				val := s.matrix[row][col]
				if min > val {
					min = val
				}
			}
			estimate += min

			// Second pass to update column minimums in the buffer
			for colIndex, col := range cols {
				if row == col {
					continue
				}
				val = s.matrix[row][col] - min

				if rowIndex == 0 {
					// First row, minimum values are set without comparison
					s.buffer[colIndex] = val
				} else {
					// Following rows, values are updated as needed
					// TODO - we can store columns with min==0 and then skip them on the next rows
					if s.buffer[colIndex] > val {
						s.buffer[colIndex] = val
					}
				}
			}
		}

		// Final pass on buffer to sum column minimums
		for colIndex := range cols {
			estimate += s.buffer[colIndex]
		}

		path := make([]types.Index, 0, len(t.Path)+1)
		path = append(path, t.Path...)
		path = append(path, nextNode)

		distance := t.Distance + s.matrix[currNode][nextNode]

		newTasks[i] = &tasks.Task{
			Path:     path,
			Distance: distance,
			Estimate: distance + estimate,
		}
	}

	return newTasks, nil
}
