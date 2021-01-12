package solver

import (
	"testing"

	"github.com/Spi1y/tsp-solver/solver/matrix"
	"github.com/Spi1y/tsp-solver/solver/tasks"
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

func runBenchamrk(b *testing.B, size int, qType tasks.QueueType) {
	b.ReportAllocs()

	matrix17 := baseMatrix17()
	sizedMatrix := matrix.ConvertToMatrix(matrix17[:size])

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := &Solver{}
		s.DistanceMatrix = sizedMatrix
		s.Solve(qType)
	}
}

func BenchmarkSolverWithList5(b *testing.B) {
	runBenchamrk(b, 5, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList6(b *testing.B) {
	runBenchamrk(b, 6, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList7(b *testing.B) {
	runBenchamrk(b, 7, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList8(b *testing.B) {
	runBenchamrk(b, 8, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList9(b *testing.B) {
	runBenchamrk(b, 9, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList10(b *testing.B) {
	runBenchamrk(b, 10, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList11(b *testing.B) {
	runBenchamrk(b, 11, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList12(b *testing.B) {
	runBenchamrk(b, 12, tasks.QueueLinkedList)
}

func BenchmarkSolverWithList13(b *testing.B) {
	runBenchamrk(b, 13, tasks.QueueLinkedList)
}

func BenchmarkSolverWithHeap5(b *testing.B) {
	runBenchamrk(b, 5, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap6(b *testing.B) {
	runBenchamrk(b, 6, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap7(b *testing.B) {
	runBenchamrk(b, 7, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap8(b *testing.B) {
	runBenchamrk(b, 8, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap9(b *testing.B) {
	runBenchamrk(b, 9, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap10(b *testing.B) {
	runBenchamrk(b, 10, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap11(b *testing.B) {
	runBenchamrk(b, 11, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap12(b *testing.B) {
	runBenchamrk(b, 12, tasks.QueueHeap)
}

func BenchmarkSolverWithHeap13(b *testing.B) {
	runBenchamrk(b, 13, tasks.QueueHeap)
}