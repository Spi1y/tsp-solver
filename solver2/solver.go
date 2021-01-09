package solver2

import "github.com/Spi1y/tsp-solver/solver2/iterator"

// Solver is a TSP solver object. It is used to set a distance matrix and start
// calculations
type Solver struct {
	i iterator.Iterator
}
