package solver

import (
	"errors"

	"github.com/Spi1y/tsp-solver/solver/matrix"
)

// Solve solves the TSP problem with a given distance matrix.
func Solve(distanceMatrix matrix.Matrix) ([]int, error) {
	size := len(distanceMatrix)

	if size == 0 {
		return nil, errors.New("empty matrix")
	}

	for _, row := range distanceMatrix {
		if len(row) != size {
			return nil, errors.New("matrix must be square")
		}
	}

	//rotNode := solutionNode{
	_ = solutionNode{
		parent:        nil,
		point:         0,
		projectedCost: 0,
		actualCost:    0,
		matrix:        nil,
	}

	return nil, nil

}
