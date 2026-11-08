package tests

import (
	"testing"

	"github.com/KelvinAkpobome/150-DSA-Questions-GO/graphs"
)

func TestHasCycleDirected(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{
			name: "cycle exists",
			n:    3,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 0},
			},
			want: true,
		},
		{
			name: "no cycle",
			n:    3,
			edges: [][]int{
				{0, 1},
				{1, 2},
			},
			want: false,
		},
		{
			name: "visited node does not mean cycle",
			n:    3,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			want: false,
		},
		{
			name: "self loop",
			n:    1,
			edges: [][]int{
				{0, 0},
			},
			want: true,
		},
		{
			name: "disconnected component contains cycle",
			n:    5,
			edges: [][]int{
				{0, 1},
				{2, 3},
				{3, 4},
				{4, 2},
			},
			want: true,
		},
		{
			name: "disconnected components without cycle",
			n:    5,
			edges: [][]int{
				{0, 1},
				{2, 3},
			},
			want: false,
		},
		{
			name:  "empty graph",
			n:     0,
			edges: [][]int{},
			want:  false,
		},
		{
			name:  "single vertex without edges",
			n:     1,
			edges: [][]int{},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graphs.HasCycleDirected(tt.n, tt.edges)

			if got != tt.want {
				t.Errorf("hasCycleDirected() = %v, want %v", got, tt.want)
			}
		})
	}
}
