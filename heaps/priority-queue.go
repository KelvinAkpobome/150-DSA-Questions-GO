package heaps

type Heap[T any] struct {
	Data []T
	Less func(a, b T) bool
}

func NewHeap[T any](Less func(a, b T) bool, items ...T) *Heap[T] {
	if Less == nil {
		panic("heap: nil comparator")
	}

	h := &Heap[T]{
		Data: append([]T(nil), items...),
		Less: Less,
	}

	for i := len(h.Data)/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}

func (h *Heap[T]) Len() int {
	return len(h.Data)
}

func (h *Heap[T]) Peek() (T, bool) {
	var zero T

	if len(h.Data) == 0 {
		return zero, false
	}

	return h.Data[0], true
}

func (h *Heap[T]) Push(val T) {
	h.Data = append(h.Data, val)
	h.up(len(h.Data) - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	var zero T

	n := len(h.Data)
	if n == 0 {
		return zero, false
	}

	top := h.Data[0]

	h.Data[0] = h.Data[n-1]

	h.Data[n-1] = zero
	h.Data = h.Data[:n-1]

	if len(h.Data) > 0 {
		h.down(0)
	}

	return top, true
}

func (h *Heap[T]) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if !h.Less(h.Data[i], h.Data[parent]) {
			break
		}

		h.Data[i], h.Data[parent] = h.Data[parent], h.Data[i]
		i = parent
	}
}

func (h *Heap[T]) down(i int) {
	n := len(h.Data)

	for {
		smallest := i
		left, right := 2*i+1, 2*i+2

		if left < n && h.Less(h.Data[left], h.Data[smallest]) {
			smallest = left
		}

		if right < n && h.Less(h.Data[right], h.Data[smallest]) {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.Data[i], h.Data[smallest] = h.Data[smallest], h.Data[i]
		i = smallest
	}
}