package main

import (
	"fmt"
	"testing"

	solver "github.com/Spi1y/tsp-solver/solver"
	solver_matrix "github.com/Spi1y/tsp-solver/solver/matrix"
	solver_tasks "github.com/Spi1y/tsp-solver/solver/tasks"

	solver2 "github.com/Spi1y/tsp-solver/solver2"
	solver2_types "github.com/Spi1y/tsp-solver/solver2/types"
)

func baseMatrix17() [][]int {
	return [][]int{
		{-1, 10000, 4940, 13206, 5941, 5866, 15147, 12730, 13714, 10632, 19693, 134984, 21742, 9385, 7930, 10139, 5281, 10263},
		{10801, -1, 11398, 6486, 5554, 6413, 21605, 4504, 3796, 3412, 14454, 129139, 17131, 4354, 5709, 3073, 6419, 295},
		{7277, 11026, -1, 15015, 5987, 5321, 11761, 12346, 17326, 14244, 18287, 131597, 20336, 9431, 7976, 10185, 4928, 11289},
		{17410, 5964, 15866, -1, 8484, 9343, 26073, 5532, 8971, 9271, 8644, 133609, 11943, 4539, 7172, 4268, 9349, 6252},
		{5550, 6532, 8185, 8658, -1, 2203, 18392, 9262, 10864, 10109, 16873, 130330, 19817, 4837, 3382, 5591, 2209, 6795},
		{4780, 5754, 7415, 8881, 1616, -1, 17622, 7975, 10086, 12531, 17117, 137459, 20061, 5060, 3605, 5814, 1863, 6017},
		{17465, 21214, 11735, 25203, 16175, 15509, -1, 22534, 27514, 24432, 28475, 130047, 30524, 19619, 18164, 20373, 15116, 21477},
		{11103, 3274, 12388, 5499, 6545, 7404, 22595, -1, 4304, 7178, 13467, 131433, 16255, 3350, 5233, 2205, 7410, 4061},
		{13688, 4936, 17473, 8938, 9338, 10197, 27680, 4628, -1, 3709, 16906, 129524, 19693, 6789, 9493, 5644, 10203, 3722},
		{10606, 3863, 14391, 9197, 8266, 9125, 24598, 6593, 3709, -1, 17165, 126442, 19842, 7065, 8421, 5784, 9131, 3931},
		{20295, 14035, 18751, 8316, 16298, 16135, 28958, 13603, 17042, 17342, -1, 148795, 4189, 11949, 14388, 10869, 17163, 14323},
		{128052, 128430, 132527, 133642, 130747, 136302, 129548, 131160, 129352, 126270, 140852, -1, 147051, 129887, 129837, 130124, 135909, 128693},
		{23594, 21180, 22050, 11615, 20041, 20900, 43627, 16165, 19604, 24462, 4189, 163225, -1, 15192, 18131, 14111, 20906, 21443},
		{9417, 3384, 9850, 4141, 4007, 4866, 20057, 3104, 6544, 6666, 11351, 130445, 13876, -1, 2695, 912, 4872, 3647},
		{6794, 3627, 8222, 7121, 2379, 3238, 18429, 4928, 7959, 6909, 15689, 130005, 18633, 2700, -1, 2937, 3244, 3890},
		{9123, 2472, 10408, 4166, 4565, 5424, 20615, 2192, 5632, 5754, 11544, 129531, 14069, 912, 3253, -1, 5430, 2735},
		{5363, 6516, 4691, 8742, 1477, 812, 14898, 7836, 10848, 9798, 16978, 134735, 19922, 4921, 3466, 5675, -1, 6779},
		{10506, 857, 11103, 6191, 5259, 6118, 21310, 3587, 4069, 3020, 14159, 128844, 16836, 4059, 5414, 2778, 6124, -1},
	}
}

func uintBaseMatrix17() [][]solver2_types.Distance {
	return [][]solver2_types.Distance{
		{0, 10000, 4940, 13206, 5941, 5866, 15147, 12730, 13714, 10632, 19693, 134984, 21742, 9385, 7930, 10139, 5281, 10263},
		{10801, 0, 11398, 6486, 5554, 6413, 21605, 4504, 3796, 3412, 14454, 129139, 17131, 4354, 5709, 3073, 6419, 295},
		{7277, 11026, 0, 15015, 5987, 5321, 11761, 12346, 17326, 14244, 18287, 131597, 20336, 9431, 7976, 10185, 4928, 11289},
		{17410, 5964, 15866, 0, 8484, 9343, 26073, 5532, 8971, 9271, 8644, 133609, 11943, 4539, 7172, 4268, 9349, 6252},
		{5550, 6532, 8185, 8658, 0, 2203, 18392, 9262, 10864, 10109, 16873, 130330, 19817, 4837, 3382, 5591, 2209, 6795},
		{4780, 5754, 7415, 8881, 1616, 0, 17622, 7975, 10086, 12531, 17117, 137459, 20061, 5060, 3605, 5814, 1863, 6017},
		{17465, 21214, 11735, 25203, 16175, 15509, 0, 22534, 27514, 24432, 28475, 130047, 30524, 19619, 18164, 20373, 15116, 21477},
		{11103, 3274, 12388, 5499, 6545, 7404, 22595, 0, 4304, 7178, 13467, 131433, 16255, 3350, 5233, 2205, 7410, 4061},
		{13688, 4936, 17473, 8938, 9338, 10197, 27680, 4628, 0, 3709, 16906, 129524, 19693, 6789, 9493, 5644, 10203, 3722},
		{10606, 3863, 14391, 9197, 8266, 9125, 24598, 6593, 3709, 0, 17165, 126442, 19842, 7065, 8421, 5784, 9131, 3931},
		{20295, 14035, 18751, 8316, 16298, 16135, 28958, 13603, 17042, 17342, 0, 148795, 4189, 11949, 14388, 10869, 17163, 14323},
		{128052, 128430, 132527, 133642, 130747, 136302, 129548, 131160, 129352, 126270, 140852, 0, 147051, 129887, 129837, 130124, 135909, 128693},
		{23594, 21180, 22050, 11615, 20041, 20900, 43627, 16165, 19604, 24462, 4189, 163225, 0, 15192, 18131, 14111, 20906, 21443},
		{9417, 3384, 9850, 4141, 4007, 4866, 20057, 3104, 6544, 6666, 11351, 130445, 13876, 0, 2695, 912, 4872, 3647},
		{6794, 3627, 8222, 7121, 2379, 3238, 18429, 4928, 7959, 6909, 15689, 130005, 18633, 2700, 0, 2937, 3244, 3890},
		{9123, 2472, 10408, 4166, 4565, 5424, 20615, 2192, 5632, 5754, 11544, 129531, 14069, 912, 3253, 0, 5430, 2735},
		{5363, 6516, 4691, 8742, 1477, 812, 14898, 7836, 10848, 9798, 16978, 134735, 19922, 4921, 3466, 5675, 0, 6779},
		{10506, 857, 11103, 6191, 5259, 6118, 21310, 3587, 4069, 3020, 14159, 128844, 16836, 4059, 5414, 2778, 6124, 0},
	}
}

func runBenchmarkCase(b *testing.B, bm bmCase) {
	b.ReportAllocs()

	switch bm.solver {
	case 1:
		matrix17 := baseMatrix17()
		sizedMatrix := solver_matrix.ConvertToMatrix(matrix17[:bm.size])
		s := &solver.Solver{}
		s.DistanceMatrix = sizedMatrix

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Solve(solver_tasks.QueueLinkedList)
		}
	case 2:
		matrix17 := baseMatrix17()
		sizedMatrix := solver_matrix.ConvertToMatrix(matrix17[:bm.size])
		s := &solver.Solver{}
		s.DistanceMatrix = sizedMatrix

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Solve(solver_tasks.QueueHeap)
		}
	case 3:
		uintMatrix17 := uintBaseMatrix17()
		sizedMatrix := uintMatrix17[:bm.size]
		for i := range sizedMatrix {
			sizedMatrix[i] = sizedMatrix[i][:bm.size]
		}
		s := &solver2.Solver{}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Solve(sizedMatrix)
		}
	}
}

func BenchmarkSolverComparison(b *testing.B) {
	benchmarks := []bmCase{
		getBMCase(5, 1),
		getBMCase(5, 2),
		getBMCase(5, 3),
		getBMCase(6, 1),
		getBMCase(6, 2),
		getBMCase(6, 3),
		getBMCase(7, 1),
		getBMCase(7, 2),
		getBMCase(7, 3),
		getBMCase(8, 1),
		getBMCase(8, 2),
		getBMCase(8, 3),
		getBMCase(9, 1),
		getBMCase(9, 2),
		getBMCase(9, 3),
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			runBenchmarkCase(b, bm)
		})
	}
}

func BenchmarkSolver1List5(b *testing.B) {
	runBenchmarkCase(b, getBMCase(5, 1))
}

func BenchmarkSolver1List7(b *testing.B) {
	runBenchmarkCase(b, getBMCase(7, 1))
}

func BenchmarkSolver1List9(b *testing.B) {
	runBenchmarkCase(b, getBMCase(9, 1))
}

func BenchmarkSolver1Heap5(b *testing.B) {
	runBenchmarkCase(b, getBMCase(5, 2))
}

func BenchmarkSolver1Heap7(b *testing.B) {
	runBenchmarkCase(b, getBMCase(7, 2))
}

func BenchmarkSolver1Heap9(b *testing.B) {
	runBenchmarkCase(b, getBMCase(9, 2))
}

func BenchmarkSolver1Heap13(b *testing.B) {
	runBenchmarkCase(b, getBMCase(13, 2))
}

func BenchmarkSolver1Heap15(b *testing.B) {
	runBenchmarkCase(b, getBMCase(15, 2))
}

func BenchmarkSolver1Heap16(b *testing.B) {
	runBenchmarkCase(b, getBMCase(16, 2))
}

func BenchmarkSolver2Heap5(b *testing.B) {
	runBenchmarkCase(b, getBMCase(5, 3))
}

func BenchmarkSolver2Heap7(b *testing.B) {
	runBenchmarkCase(b, getBMCase(7, 3))
}

func BenchmarkSolver2Heap9(b *testing.B) {
	runBenchmarkCase(b, getBMCase(9, 3))
}

func BenchmarkSolver2Heap11(b *testing.B) {
	runBenchmarkCase(b, getBMCase(11, 3))
}

func BenchmarkSolver2Heap13(b *testing.B) {
	runBenchmarkCase(b, getBMCase(13, 3))
}

func BenchmarkSolver2Heap15(b *testing.B) {
	runBenchmarkCase(b, getBMCase(15, 3))
}

func BenchmarkSolver2Heap16(b *testing.B) {
	runBenchmarkCase(b, getBMCase(16, 3))
}

func solverString(s int) string {
	switch s {
	case 1:
		return "Solver1	List"
	case 2:
		return "Solver1	Heap"
	case 3:
		return "Solver2	Heap"
	default:
		return "Unknown"
	}
}

type bmCase struct {
	name   string
	size   int
	solver int
}

func getBMCase(size int, solver int) bmCase {
	return bmCase{
		name: fmt.Sprintf("%v	%v", size, solverString(solver)),
		size:   size,
		solver: solver,
	}
}
