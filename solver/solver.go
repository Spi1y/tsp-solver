package solver

import "errors"

// Solve solves the TSP problem with a given distance matrix.
// Distance matrix must be square.
func Solve(distanceMatrix [][]int) ([]int, error) {
	size := len(distanceMatrix)

	if size == 0 {
		return nil, errors.New("empty matrix")
	}

	for _, row := range distanceMatrix {
		if len(row) != size {
			return nil, errors.New("matrix must be square")
		}
	}

	return nil, nil

}
