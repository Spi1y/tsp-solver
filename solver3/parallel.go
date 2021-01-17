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

type processingPacket struct {
	task     tasks.Task
	newTasks []tasks.Task
	solution struct {
		path     []types.Index
		distance types.Distance
	}
}

func (s *Solver) taskProcessor(in <-chan *processingPacket, out chan<- *processingPacket, done <-chan struct{}) {
	size := len(s.matrix)
	iter := &iterator.Iterator{}
	iter.Init(types.Index(size))
	buff := make([]types.Distance, size)

	for {
		select {
		case pkt := <-in:
			err := s.processTask(iter, buff, pkt)
			if err != nil {
				panic(err)
			}
			out <- pkt
		case <-done:
			return
		}
	}
}

func (s *Solver) processTask(it *iterator.Iterator, buf []types.Distance, pkt *processingPacket) error {
	// TODO - try aggressive approach with full path first

	t := pkt.task
	err := it.SetPath(t.Path)
	if err != nil {
		return err
	}
	nextNodes := it.NodesToVisit()
	rows := it.RowsToIterate()

	currNode := t.Path[len(t.Path)-1]
	nodesLeft := len(nextNodes)

	if nodesLeft <= int(s.RecursiveThreshold) {
		// Calculate remaining path through brute-force recursion
		tailpath, taildistance := s.solveRecursively(currNode, nextNodes)
		path := make([]types.Index, len(t.Path), len(t.Path)+len(tailpath))
		copy(path, t.Path)

		pkt.solution.path = append(path, tailpath...)
		pkt.solution.distance = t.Distance + taildistance
		pkt.newTasks = pkt.newTasks[:0]

		return nil
	}

	if nodesLeft == 1 {
		// Final node, calculate return distance to root node and publish solution
		finalNode := nextNodes[0]

		path := make([]types.Index, len(t.Path), len(t.Path)+2)
		copy(path, t.Path)

		pkt.solution.path = append(path, finalNode, 0)
		pkt.solution.distance = t.Distance + s.matrix[currNode][finalNode] + s.matrix[finalNode][0]
		pkt.newTasks = pkt.newTasks[:0]
		return nil
	}

	newPathLen := len(t.Path) + 1
	pathsSlice := make([]types.Index, nodesLeft*newPathLen)
	pkt.newTasks = pkt.newTasks[:nodesLeft]
	pkt.solution.path = pkt.solution.path[:0]

	for i, nextNode := range nextNodes {

		var estimate types.Distance
		cols, err := it.ColsToIterate(nextNode)
		if err != nil {
			return err
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

		pkt.newTasks[i].Path = path
		pkt.newTasks[i].Distance = distance
		pkt.newTasks[i].Estimate = distance + estimate
	}

	return nil
}

func (s *Solver) newPacket() *processingPacket {
	pkt := processingPacket{}
	pkt.newTasks = make([]tasks.Task, len(s.matrix))
	pkt.solution.path = make([]types.Index, 0, len(s.matrix)+1)

	return &pkt
}

func (s *Solver) solveParallel() {
	if s.taskQueue.IsEmpty() {
		return
	}

	threadscount := runtime.NumCPU() - 1
	if threadscount == 0 {
		threadscount = 1
	}

	toProcessors := make(chan *processingPacket, threadscount)
	fromProcessors := make(chan *processingPacket, threadscount)
	stopProcessing := make(chan struct{}, threadscount)

	for i := 0; i < threadscount; i++ {
		go s.taskProcessor(toProcessors, fromProcessors, stopProcessing)
	}

	var pkt *processingPacket
	busyThreads := 0

	// Sending initial tasks
	for i := 0; i < threadscount; i++ {
		if s.taskQueue.IsEmpty() {
			break
		}
		// And we have new work for them
		task, err := s.taskQueue.PopFirst()
		if err != nil {
			panic(err)
		}

		pkt := s.newPacket()
		pkt.task = task
		toProcessors <- pkt
		busyThreads++
	}

	for {
		select {
		case pkt = <-fromProcessors:
			if len(pkt.newTasks) != 0 {
				s.taskQueue.Insert(pkt.newTasks)
			}
			if len(pkt.solution.path) != 0 {
				s.newSolutionFound(pkt.solution.path, pkt.solution.distance)
			}

			if !s.taskQueue.IsEmpty() {
				task, err := s.taskQueue.PopFirst()
				if err != nil {
					panic(err)
				}

				pkt.task = task
				toProcessors <- pkt
			} else {
				busyThreads--
			}
		}

		for busyThreads != threadscount && !s.taskQueue.IsEmpty() {
			// There are processors waiting for work
			// and we have new work for them
			task, err := s.taskQueue.PopFirst()
			if err != nil {
				panic(err)
			}

			pkt = s.newPacket()
			pkt.task = task
			toProcessors <- pkt
			busyThreads++
		}

		if s.taskQueue.IsEmpty() && busyThreads == 0 {
			break
		}
	}

	for i := 0; i < threadscount; i++ {
		stopProcessing <- struct{}{}
	}
}
