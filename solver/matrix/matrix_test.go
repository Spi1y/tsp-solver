package matrix

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToMatrix(t *testing.T) {
	type args struct {
		slice [][]int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			"Empty",
			args{
				slice: [][]int{},
			},
			Matrix{},
		},
		{
			"Normal - 3*3",
			args{
				slice: [][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
			Matrix{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
		},
		{
			"Special - 3*2",
			args{
				slice: [][]int{
					{1, 2},
					{1, 2},
					{1, 2},
				},
			},
			Matrix{
				{1, 2, 0},
				{1, 2, 0},
				{1, 2, 0},
			},
		},
		{
			"Special - 3*4",
			args{
				slice: [][]int{
					{1, 2, 3, 4},
					{1, 2, 3, 4},
					{1, 2, 3, 4},
				},
			},
			Matrix{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToMatrix(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Matrix_copy(t *testing.T) {
	emptyMatrix := ConvertToMatrix(nil)
	slice2_2 := [][]int{
		{0, 1},
		{1, 0},
	}
	Matrix2_2 := ConvertToMatrix(slice2_2)

	tests := []struct {
		name string
		m    Matrix
		want Matrix
	}{
		{
			"empty",
			emptyMatrix,
			emptyMatrix,
		},
		{
			"2*2",
			Matrix2_2,
			Matrix2_2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Copy()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.copy() = %v, want %v", got, tt.want)
			}

			if len(got) > 0 {
				if &got[0] == &tt.m[0] {
					t.Errorf("Matrix.copy() should not be a ref to an argument")
				}
			}
		})
	}
}

func TestMatrix_Normalize(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want int
		mOut Matrix
	}{
		{
			"empty",
			ConvertToMatrix(nil),
			0,
			ConvertToMatrix(nil),
		},
		{
			"2*2",
			ConvertToMatrix([][]int{
				{1, 5},
				{4, 9},
			}),
			9,
			ConvertToMatrix([][]int{
				{0, 0},
				{0, 1},
			}),
		},
		{
			"4*4",
			ConvertToMatrix([][]int{
				{1, 5, 5, 2},
				{4, 9, 5, 1},
				{1, 5, 4, 2},
				{4, 9, 8, 6},
			}),
			14,
			ConvertToMatrix([][]int{
				{0, 0, 1, 1},
				{3, 4, 1, 0},
				{0, 0, 0, 1},
				{0, 1, 1, 2},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Normalize()
			if got != tt.want {
				t.Errorf("Matrix.Normalize() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.mOut, tt.m) {
				t.Errorf("After Matrix.Normalize() m became %v, should be %v", tt.m, tt.mOut)
			}
		})
	}
}

func TestMatrix_CutNode(t *testing.T) {
	type args struct {
		source   int
		dest     int
		lastNode bool
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		mOut Matrix
	}{
		{
			"empty",
			ConvertToMatrix(nil),
			args{
				source:   0,
				dest:     0,
				lastNode: false,
			},
			ConvertToMatrix(nil),
		},
		{
			"3*3",
			ConvertToMatrix([][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 2, 1},
			}),
			args{
				source:   0,
				dest:     1,
				lastNode: false,
			},
			ConvertToMatrix([][]int{
				{-1, -1, -1},
				{-1, -1, 1},
				{3, -1, 1},
			}),
		},
		{
			"3*3, last node",
			ConvertToMatrix([][]int{
				{-1, -1, -1},
				{-1, -1, 1},
				{3, -1, 1},
			}),
			args{
				source:   1,
				dest:     2,
				lastNode: true,
			},
			ConvertToMatrix([][]int{
				{-1, -1, -1},
				{-1, -1, -1},
				{3, -1, -1},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.CutNode(tt.args.source, tt.args.dest, tt.args.lastNode)

			if !reflect.DeepEqual(tt.mOut, tt.m) {
				t.Errorf("After Matrix.CutNode() m became %v, should be %v", tt.m, tt.mOut)
			}
		})
	}
}

func TestMatrix_LoadFrom(t *testing.T) {
	tests := []struct {
		name    string
		m       Matrix
		source  Matrix
		mOut    Matrix
		wantErr bool
	}{
		{
			"empty",
			ConvertToMatrix(nil),
			ConvertToMatrix(nil),
			ConvertToMatrix(nil),
			false,
		},
		{
			"3*3",
			ConvertToMatrix([][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			}),
			ConvertToMatrix([][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 2, 1},
			}),
			ConvertToMatrix([][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 2, 1},
			}),
			false,
		},
		{
			"Size mismatch",
			ConvertToMatrix([][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 2, 1},
			}),
			ConvertToMatrix([][]int{
				{1, 2, 3, 4},
				{2, 3, 1, 4},
				{3, 2, 1, 4},
				{3, 2, 1, 4},
			}),
			ConvertToMatrix([][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 2, 1},
			}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.m.LoadFrom(tt.source)
			assert.Equal(t, tt.mOut, tt.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
