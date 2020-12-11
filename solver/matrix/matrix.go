package matrix

import "errors"

// Matrix is a square matrix.
// It`s underlying [][]int slice is guaranteed to be sliced from one
// linear backing array, which allows some copying optimizations.
type Matrix [][]int

// ConvertToMatrix converts raw [][]int slice to square matrix
func ConvertToMatrix(slice [][]int) Matrix {
	size := len(slice)

	backingArray := make([]int, size*size)
	var matrix Matrix = make([][]int, size)

	for i, row := range slice {
		matrix[i] = backingArray[i*size : (i+1)*size]
		copy(matrix[i], row)
	}

	return matrix
}

// Copy copies matrix to a new one with some optimizations.
func (m Matrix) Copy() Matrix {
	size := len(m)

	backingArray := make([]int, size*size)
	if size != 0 {
		copy(backingArray, m[0][:cap(m[0])])
	}

	var matrix Matrix = make([][]int, size)

	for i := range m {
		matrix[i] = backingArray[i*size : (i+1)*size]
	}

	return matrix
}

// LoadFrom copies matrix data from another matrix
func (m Matrix) LoadFrom(source Matrix) error {
	size := len(m)

	if len(m) != len(source) {
		return errors.New("Matrix size mismatch")
	}

	if size != 0 {
		copy(m[0][:cap(m[0])], source[0][:cap(m[0])])
	}

	return nil
}

// Normalize calculates minimal value in each row separately.
// Then it substracts this value from other values in the row.
// Then, a process is repeated for all columns.
// All minimum values are summed and returned as a return value.
func (m Matrix) Normalize() int {
	size := len(m)
	norm := 0
	val := 0

	for i := 0; i < size; i++ {
		// First, we calculate minimum
		min := -1
		for j := 0; j < size; j++ {
			val = m[i][j]

			if val == -1 {
				continue
			}

			if val == 0 {
				min = 0
				break
			}

			if (min == -1) || (min > val) {
				min = val
			}
		}

		if min <= 0 {
			continue
		}

		// And now, we substract it
		for j := 0; j < size; j++ {
			val = m[i][j]
			if val == -1 {
				continue
			}

			m[i][j] = val - min
		}

		// Accumulation
		norm += min
	}

	// And we repeat the same process for columns
	// I couldn`t figure out a way to get rid of duplication yet
	for j := 0; j < size; j++ {
		// First, we calculate minimum
		min := -1
		for i := 0; i < size; i++ {
			val = m[i][j]

			if val == -1 {
				continue
			}

			if val == 0 {
				min = 0
				break
			}

			if (min == -1) || (min > val) {
				min = val
			}
		}

		if min <= 0 {
			continue
		}

		// And now, we substract it
		for i := 0; i < size; i++ {
			val = m[i][j]
			if val == -1 {
				continue
			}

			m[i][j] = val - min
		}

		// Accumulation
		norm += min
	}

	return norm
}

// CutNode cuts elements from matrix (sets them to -1) according
// to the processed path from source to destination nodes.
// Details included in the method body.
func (m Matrix) CutNode(source, dest int, lastNode bool) {
	size := len(m)

	// Disabling paths from the source - we have already passed it
	for i := 0; i < size; i++ {
		m[source][i] = -1
	}

	// Disabling paths to the destination - we are already in it
	for i := 0; i < size; i++ {
		m[i][dest] = -1
	}

	if !lastNode && (size != 0) {
		// Disabling path to the root node (index 0).
		// There are more nodes to visit, and we will be coming to
		// root from one of them.
		m[dest][0] = -1
	}
}
