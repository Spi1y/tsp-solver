package solver

import "github.com/Spi1y/tsp-solver/solver/matrix"

type solutionNode struct {
	parent        *solutionNode
	point         int
	projectedCost int
	actualCost    int

	matrix matrix.Matrix
}

func newRootNode(matrix matrix.Matrix) *solutionNode {

	return &solutionNode{
		parent:        nil,
		point:         0,
		projectedCost: 0,
		actualCost:    0,
		matrix:        nil,
	}
}
