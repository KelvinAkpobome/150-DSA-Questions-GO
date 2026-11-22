package tests

import (
	"testing"

	"github.com/KelvinAkpobome/150-DSA-Questions-GO/graphs"
)

func TestNumberOfIslands(t *testing.T) {
	tests := []struct {
		name string
		want int8
	}{
		{
			name: "three disconnected islands",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graphs.NumberOfIslands()

			if got != tt.want {
				t.Errorf("NumberOfIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}
