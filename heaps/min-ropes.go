package heaps

func ConnectRopes(ropes []int) int {
    h := NewHeap(func(a, b int) bool {
        return a < b
    }, ropes...)

    cost := 0

    for h.Len() >= 2 {
        f, _ := h.Pop()
        s, _ := h.Pop()

        combined := f + s

        cost += combined
        h.Push(combined)
    }

    return cost
}

