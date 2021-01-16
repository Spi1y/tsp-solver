package tasks

import "github.com/Spi1y/tsp-solver/solver2/types"

// Task represents a step of the algorithm
// It includes all data necessary for the calculation of a step
// and generation of next steps
type Task struct {
	// An ordered list of the nodes, representing already traveled path
	Path []types.Index
	// Distance traveled while traversing the CurrentPath
	Distance types.Distance
	// Estimate of lowest possible distance on that path
	// Used for prioritization of tasks
	Estimate types.Distance
}
