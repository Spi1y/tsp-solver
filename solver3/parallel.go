package solver3

import (
	"runtime"

	"github.com/Spi1y/tsp-solver/solver2/iterator"
	"github.com/Spi1y/tsp-solver/solver2/tasks"
	"github.com/Spi1y/tsp-solver/solver2/types"
)

type solution struct {
	path     []types.Index
	distance types.Distance
}

func (s *Solver) taskProcessor(ready chan<- struct{}, task <-chan tasks.Task, results chan<- tasks.Task, solutions chan<- solution, done <-chan struct{}) {
	ready <- struct{}{}
	newTasks := make([]tasks.Task, len(s.matrix))

	size := len(s.matrix)
	it := &iterator.Iterator{}
	it.Init(types.Index(size))
	buf := make([]types.Distance, size)

	for {
		select {
		case t := <-task:
			taskcount, solution, err := s.processTask(it, buf, t, newTasks)
			if err != nil {
				panic(err)
			}
			if solution != nil {
				solutions <- *solution
			}
			for _, t := range newTasks[:taskcount] {
				results <- t
			}
			ready <- struct{}{}
		case <-done:
			return
		}
	}
}

func (s *Solver) processTask(it *iterator.Iterator, buf []types.Distance, t tasks.Task, newTasks []tasks.Task) (int, *solution, error) {
	// TODO - try aggressive approach with full path first

	err := it.SetPath(t.Path)
	if err != nil {
		return 0, nil, err
	}
	nextNodes := it.NodesToVisit()
	rows := it.RowsToIterate()

	currNode := t.Path[len(t.Path)-1]
	nodesLeft := len(nextNodes)

	if nodesLeft <= int(s.RecursiveThreshold) {
		tailpath, taildistance := s.solveRecursively(currNode, nextNodes)
		path := make([]types.Index, len(t.Path), len(t.Path)+len(tailpath))
		copy(path, t.Path)
		path = append(path, tailpath...)
		distance := t.Distance + taildistance

		newTasks = newTasks[:0]
		return 0, &solution{path: path, distance: distance}, nil
	}

	if nodesLeft == 1 {
		// Final node, calculating return distance to root node
		// and notifying solver about found solution
		finalNode := nextNodes[0]

		path := make([]types.Index, len(t.Path), len(t.Path)+2)
		copy(path, t.Path)
		path = append(path, finalNode, 0)
		distance := t.Distance + s.matrix[currNode][finalNode] + s.matrix[finalNode][0]

		newTasks = newTasks[:0]
		return 0, &solution{path: path, distance: distance}, nil
	}

	newPathLen := len(t.Path) + 1
	pathsSlice := make([]types.Index, nodesLeft*newPathLen)

	for i, nextNode := range nextNodes {

		var estimate types.Distance
		cols, err := it.ColsToIterate(nextNode)
		if err != nil {
			return 0, nil, err
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

				val = rowSlice[col]
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
					buf[colIndex] = rowSlice[col] - min
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
				if buf[colIndex] > val {
					buf[colIndex] = val
				}
			}
		}

		// Final pass on buffer to sum column minimums
		for colIndex := range cols {
			estimate += buf[colIndex]
		}

		path := pathsSlice[i*newPathLen : (i+1)*newPathLen]
		copy(path, t.Path)
		path[newPathLen-1] = nextNode

		distance := t.Distance + s.matrix[currNode][nextNode]

		newTasks[i].Path = path
		newTasks[i].Distance = distance
		newTasks[i].Estimate = distance + estimate
	}

	return nodesLeft, nil, nil
}

func (s *Solver) solveParallel() {
	if s.taskQueue.IsEmpty() {
		return
	}

	threadscount := runtime.NumCPU() - 1
	if threadscount == 0 {
		threadscount = 1
	}

	readyToProcess := make(chan struct{}, threadscount)
	tasksToProcess := make(chan tasks.Task, threadscount)
	tasksToQueue := make(chan tasks.Task, threadscount*len(s.matrix))
	solutionFound := make(chan solution, threadscount)
	stopProcessing := make(chan struct{}, threadscount)

	for i := 0; i < threadscount; i++ {
		go s.taskProcessor(readyToProcess, tasksToProcess, tasksToQueue, solutionFound, stopProcessing)
	}

	busyThreads := threadscount

	var noTasksLeft, chanellsAreClear bool
	for {
		select {
		case <-readyToProcess:
			busyThreads--
		case taskToQueue := <-tasksToQueue:
			s.taskQueue.InsertSingle(taskToQueue)
		case solution := <-solutionFound:
			s.newSolutionFound(solution.path, solution.distance)
		}

		if threadscount != busyThreads {
			task, err := s.taskQueue.PopFirst()
			if err == nil {
				tasksToProcess <- task
				busyThreads++
			}
		}

		noTasksLeft = s.taskQueue.IsEmpty() && busyThreads == 0
		chanellsAreClear = len(tasksToQueue) == 0 && len(solutionFound) == 0
		if noTasksLeft && chanellsAreClear {
			break
		}
	}

	for i := 0; i < threadscount; i++ {
		stopProcessing <- struct{}{}
	}
}
