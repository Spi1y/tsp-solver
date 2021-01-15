package solver2

import (
	"runtime"

	"github.com/Spi1y/tsp-solver/solver2/tasks"
	"github.com/Spi1y/tsp-solver/solver2/types"
)

type solution struct {
	path     []types.Index
	distance types.Distance
}

func (s *Solver) taskProcessor(ready chan<- struct{}, task <-chan tasks.Task, results chan<- tasks.Task, solutions chan<- solution, done <-chan struct{}) {

}

func (s *Solver) queueProcessor(done chan<- struct{}) {
	if s.taskQueue.IsEmpty() {
		done <- struct{}{}
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

	for !(s.taskQueue.IsEmpty() && busyThreads == 0) {
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
	}
}
