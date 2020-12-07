package tasks

import "github.com/Spi1y/tsp-solver/solver/matrix"

// Task represents a step of the algorithm
// It includes all data necessary for the calculation of a step
//  and generation of next steps
type Task struct {
	// An ordered list of the nodes, representing already traveled path
	CurrentPath []int
	// Distance traveled while traversing the CurrentPath
	ActualDistance int

	// Next node of the path, the one we plan to advance to
	NextNode int
	// Estimation of the lower-bound of the distance of all possible paths
	// through NextNode to the end. Calculated in the Matrix.Normalize().
	ProjectedDistance int

	// Distance matrix transformed according to CurrentPath
	DistanceMatrix matrix.Matrix
}
