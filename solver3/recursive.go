package solver3

import (
	"github.com/Spi1y/tsp-solver/solver2/types"
)

// solveRecursively solves TSP recursively (brute-force). Depite its exponential O(),
// on small matrices it will be faster than other "smarter" algorithms
func (s *Solver) solveRecursively(currNode types.Index, nextNodes []types.Index) ([]types.Index, types.Distance) {
	bestPath := make([]types.Index, 0, len(nextNodes)+2)
	var bestDistance types.Distance

	permutationProcessor := func(path []types.Index) {
		dist := s.calculatePath(currNode, path)
		if len(bestPath) == 0 {
			bestDistance = dist
			bestPath = append(bestPath, path...)
			bestPath = append(bestPath, 0)
		} else if bestDistance > dist {
			bestDistance = dist
			copy(bestPath, path)
		}
	}

	calculatePermutations(nextNodes, permutationProcessor, 0)

	return bestPath, bestDistance
}

func calculatePermutations(nodes []types.Index, fn func([]types.Index), index int) {
	if index > len(nodes) {
		fn(nodes)
		return
	}

	calculatePermutations(nodes, fn, index+1)
	for j := index + 1; j < len(nodes); j++ {
		nodes[index], nodes[j] = nodes[j], nodes[index]
		calculatePermutations(nodes, fn, index+1)
		nodes[index], nodes[j] = nodes[j], nodes[index]
	}
}

func (s *Solver) calculatePath(currNode types.Index, path []types.Index) types.Distance {
	l := len(path)

	dist := s.matrix[currNode][path[0]]
	for i := 0; i < l-1; i++ {
		dist += s.matrix[path[i]][path[i+1]]
	}
	dist += s.matrix[path[l-1]][0]

	return dist
}
