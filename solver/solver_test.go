package solver

import (
	"reflect"
	"testing"

	"github.com/Spi1y/tsp-solver/solver/matrix"
)

func TestSolve(t *testing.T) {
	type args struct {
		distanceMatrix matrix.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"nil",
			args{
				nil,
			},
			nil,
			true,
		},
		{
			"non square",
			args{
				matrix.Matrix{
					{0, 1},
					{0, 1},
					{0, 1},
				},
			},
			nil,
			true,
		},
		{
			"non square 2",
			args{
				matrix.Matrix{
					{0, 1},
					{0},
				},
			},
			nil,
			true,
		},
		{
			"problem - 2",
			args{
				matrix.Matrix{
					{0, 1, 9},
					{9, 0, 1},
					{1, 9, 0},
				},
			},
			[]int{0, 1, 2, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.args.distanceMatrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
