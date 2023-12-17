package aoc2023

type Heap[T any] struct {
	s       []T
	compare func(a, b T) int
}

func NewHeap[T any](compare func(a, b T) int) *Heap[T] {
	return &Heap[T]{compare: compare}
}

func (h *Heap[T]) Push(v T) {
	i := len(h.s)
	h.s = append(h.s, v)
	for {
		parent := (i - 1) / 2
		if parent == i || h.compare(h.s[parent], v) >= 0 {
			break
		}
		h.s[parent], h.s[i] = h.s[i], h.s[parent]
		i = parent
	}
}

func (h *Heap[T]) Pop() T {
	if len(h.s) == 0 {
		panic("empty heap")
	}
	v := h.s[0]
	h.s[0] = h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]
	i := 0
	for {
		l := i*2 + 1
		if l >= len(h.s) {
			break
		}
		r := l + 1
		child := l
		if r < len(h.s) && h.compare(h.s[r], h.s[l]) > 0 {
			child = r
		}
		if h.compare(h.s[child], h.s[i]) <= 0 {
			break
		}
		h.s[child], h.s[i] = h.s[i], h.s[child]
		i = child
	}
	return v
}

func (h *Heap[T]) Len() int {
	return len(h.s)
}
