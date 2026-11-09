package tests

import (
	"reflect"
	"testing"

	"github.com/KelvinAkpobome/150-DSA-Questions-GO/graphs"
)

func TestDFSRecursive(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[int][]int
		start    int
		expected []int
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
			start:    1,
			expected: []int{1, 2, 4, 5, 3, 6},
		},
		{
			name: "single node",
			graph: map[int][]int{
				1: {},
			},
			start:    1,
			expected: []int{1},
		},
		{
			name: "linear graph",
			graph: map[int][]int{
				1: {2},
				2: {1, 3},
				3: {2, 4},
				4: {3},
			},
			start:    1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "cyclic graph",
			graph: map[int][]int{
				1: {2, 3},
				2: {1, 3},
				3: {1, 2},
			},
			start:    1,
			expected: []int{1, 2, 3},
		},
		{
			name: "disconnected graph",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
			},
			start:    1,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visited := make(map[int]bool)
			order := make([]int, 0)

			got := graphs.DFSRecursive(
				tt.graph,
				tt.start,
				visited,
				order,
			)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("DFSRecursive() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDFSIterative(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[int][]int
		start    int
		expected []int
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
			start:    1,
			expected: []int{1, 3, 6, 5, 4, 2},
		},
		{
			name: "single node",
			graph: map[int][]int{
				1: {},
			},
			start:    1,
			expected: []int{1},
		},
		{
			name: "linear graph",
			graph: map[int][]int{
				1: {2},
				2: {1, 3},
				3: {2, 4},
				4: {3},
			},
			start:    1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "cyclic graph",
			graph: map[int][]int{
				1: {2, 3},
				2: {1, 3},
				3: {1, 2},
			},
			start:    1,
			expected: []int{1, 3, 2},
		},
		{
			name: "disconnected graph",
			graph: map[int][]int{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
			},
			start:    1,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graphs.DFSIterative(tt.graph, tt.start)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("DFSIterative() = %v, want %v", got, tt.expected)
			}
		})
	}
}