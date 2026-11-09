package tests

import (
	"reflect"
	"testing"

	"github.com/KelvinAkpobome/150-DSA-Questions-GO/graphs"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name  string
		graph map[int][]int
		start int
		want  []int
	}{
		{
			name: "connected graph",
			graph: map[int][]int{
				1: {2, 3},
				2: {1, 4},
				3: {1, 5, 6},
				4: {2, 5},
				5: {3, 4},
				6: {3},
			},
			start: 1,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "single node",
			graph: map[int][]int{
				1: {},
			},
			start: 1,
			want:  []int{1},
		},
		{
			name: "linear graph",
			graph: map[int][]int{
				1: {2},
				2: {1, 3},
				3: {2, 4},
				4: {3},
			},
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name: "cyclic graph",
			graph: map[int][]int{
				1: {2, 3},
				2: {1, 3},
				3: {1, 2},
			},
			start: 1,
			want:  []int{1, 2, 3},
		},
		{
			name: "disconnected graph",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
			},
			start: 1,
			want:  []int{1, 2},
		},
		{
			name: "start node has no neighbours",
			graph: map[int][]int{
				1: {},
				2: {3},
				3: {2},
			},
			start: 1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graphs.BFS(tt.graph, tt.start)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}