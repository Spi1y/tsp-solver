package tasks

import "github.com/Spi1y/tsp-solver/solver/matrix"

// Task represents a step of the algorithm
// It includes all data necessary for the calculation of a step
// and generation of next steps
type Task struct {
	// An ordered list of the nodes, representing already traveled path
	Path []int
	// Distance traveled while traversing the CurrentPath
	Distance int

	// Distance matrix transformed according to CurrentPath
	DistanceMatrix matrix.Matrix
}
