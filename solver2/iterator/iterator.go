package iterator

import (
	"fmt"

	"github.com/Spi1y/tsp-solver/solver2/types"
)

// Iterator is a calculator used to determine the following values:
// - nodes left to visit (based on currently traveled path)
// - rows and columns used to calculate distance lower estimate (based on the path and the next node)
type Iterator struct {
	// Size of the matrix, set in the Init()
	size types.Index

	// Mask slice of visited nodes, used to simplify path processing
	nodesVisited []bool

	// Last known node of the path
	lastNode int16

	// Calculated lists
	nodesToVisit  []types.Index
	colsToIterate []types.Index
	//rowsToIterate []types.Index
}

// Init is used to initialize internal structures according to the distance
// matrix size
func (it *Iterator) Init(size types.Index) {
	it.size = size

	it.nodesVisited = make([]bool, size)
	it.nodesToVisit = make([]types.Index, size)
	it.colsToIterate = make([]types.Index, size)
	//it.rowsToIterate = make([]types.Index, size)
}

// SetPath process given path and calculates internal data structures
func (it *Iterator) SetPath(path []types.Index) error {
	if it.size == 0 {
		return fmt.Errorf("Iterator is not initialized")
	}

	if (len(path) == 0) || (path[0] != 0) {
		return fmt.Errorf("Path must include 0 node as the first element")
	}

	if it.size < types.Index(len(path)) {
		return fmt.Errorf("Incorrect path: length %v is greater than matrix size %v", len(path), it.size)
	}

	nodesCount := it.size - types.Index(len(path))
	it.nodesToVisit = it.nodesToVisit[:nodesCount]

	it.resetBuf()

	for _, node := range path {
		if node >= it.size {
			return fmt.Errorf("Wrong node in the path: index %v is greater than matrix size %v", node, it.size)
		}
		it.nodesVisited[node] = true
	}

	if it.size == types.Index(len(path)) {
		return nil
	}

	c := 0
	for i := 0; types.Index(i) < it.size; i++ {
		if it.nodesVisited[i] == false {
			it.nodesToVisit[c] = types.Index(i)
			c++
		}
	}

	if len(path) == 0 {
		it.lastNode = -1
	} else {
		it.lastNode = int16(path[len(path)-1])
	}

	return nil
}

// NodesToVisit retrieves the list of nodes left to visit
func (it *Iterator) NodesToVisit() []types.Index {
	return it.nodesToVisit
}

// ColsToIterate is used to calculate the list of column indices which have to
// be processed to determine distance lower estimate, based on the path and the next node
func (it *Iterator) ColsToIterate(node types.Index) ([]types.Index, error) {
	if node >= it.size {
		return nil, fmt.Errorf("Incorrect next node index %v", node)
	}

	if len(it.nodesToVisit) == 0 {
		if node != 0 {
			return nil, fmt.Errorf("Incorrect next node index %v", node)
		}

		it.colsToIterate = it.colsToIterate[:0]
		return it.colsToIterate, nil
	}

	if it.nodesVisited[node] == true {
		return nil, fmt.Errorf("Node %v already visited", node)
	}

	var index types.Index
	for i, val := range it.nodesToVisit {
		if val == node {
			index = types.Index(i)
			break
		}
	}

	// We do not decrement len because we need one additional element to hold 0 node
	it.colsToIterate = it.colsToIterate[:len(it.nodesToVisit)]

	it.colsToIterate[0] = 0
	copy(it.colsToIterate[1:index+1], it.nodesToVisit[:index])
	copy(it.colsToIterate[index+1:], it.nodesToVisit[index+1:])

	return it.colsToIterate, nil
}

// RowsToIterate is used to calculate the list of row indices which have to
// be processed to determine distance lower estimate, based on the path and the next node
// It is equal to it.nodesToVisit
func (it *Iterator) RowsToIterate() []types.Index {
	return it.nodesToVisit
}

// resetBuf resets values in the buffer to false using copy optimization
func (it *Iterator) resetBuf() {
	if len(it.nodesVisited) == 0 {
		return
	}

	it.nodesVisited[0] = false
	for j := 1; types.Index(j) < it.size; j *= 2 {
		copy(it.nodesVisited[j:], it.nodesVisited[:j])
	}
}
