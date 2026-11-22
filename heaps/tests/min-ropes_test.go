package tests

import (
	"testing"

	"github.com/KelvinAkpobome/150-DSA-Questions-GO/heaps"
)

func TestConnectRopes(t *testing.T) {
	tests := []struct {
		name  string
		ropes []int
		want  int
	}{
		{
			name:  "multiple ropes",
			ropes: []int{4, 3, 2, 6},
			want:  29,
		},
		{
			name:  "two ropes",
			ropes: []int{4, 3},
			want:  7,
		},
		{
			name:  "single rope",
			ropes: []int{10},
			want:  0,
		},
		{
			name:  "empty input",
			ropes: []int{},
			want:  0,
		},
		{
			name:  "all equal lengths",
			ropes: []int{5, 5, 5, 5},
			want:  40,
		},
		{
			name:  "already sorted",
			ropes: []int{1, 2, 3, 4, 5},
			want:  33,
		},
		{
			name:  "unsorted ropes",
			ropes: []int{8, 4, 6, 12},
			want:  58,
		},
		{
			name:  "ropes with zero length",
			ropes: []int{0, 0, 5, 10},
			want:  20,
		},
		{
			name:  "large rope lengths",
			ropes: []int{100, 200, 300},
			want:  900,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := heaps.ConnectRopes(tt.ropes)

			if got != tt.want {
				t.Errorf("ConnectRopes(%v) = %d, want %d",
					tt.ropes, got, tt.want)
			}
		})
	}
}


func BenchmarkHeapPushPop(b *testing.B) {
    h := &heaps.Heap[int]{
        Data: make([]int, 5, 10),
        Less: func(a, b int) bool {
            return a < b
        },
    }

    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        h.Push(10)
        h.Pop()
    }
}