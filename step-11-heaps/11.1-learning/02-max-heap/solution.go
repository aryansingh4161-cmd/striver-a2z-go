package maxheap

import "container/heap"

// MaxHeap is a hand-written binary max-heap.
type MaxHeap struct{ a []int }

func (h *MaxHeap) Len() int { return len(h.a) }

// Push inserts x. O(log n).
func (h *MaxHeap) Push(x int) {
	h.a = append(h.a, x)
	i := len(h.a) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.a[p] >= h.a[i] {
			break
		}
		h.a[p], h.a[i] = h.a[i], h.a[p]
		i = p
	}
}

// Pop removes and returns the maximum. O(log n). Panics if empty.
func (h *MaxHeap) Pop() int {
	top := h.a[0]
	last := len(h.a) - 1
	h.a[0] = h.a[last]
	h.a = h.a[:last]
	i, n := 0, len(h.a)
	for {
		big := i
		l, r := 2*i+1, 2*i+2
		if l < n && h.a[l] > h.a[big] {
			big = l
		}
		if r < n && h.a[r] > h.a[big] {
			big = r
		}
		if big == i {
			break
		}
		h.a[i], h.a[big] = h.a[big], h.a[i]
		i = big
	}
	return top
}

// IntMaxHeap is the container/heap version: only Less is flipped.
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Drain pops everything from a container/heap max-heap (helper for tests/demo).
func Drain(values []int) []int {
	h := &IntMaxHeap{}
	for _, v := range values {
		heap.Push(h, v)
	}
	out := []int{}
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(int))
	}
	return out
}
