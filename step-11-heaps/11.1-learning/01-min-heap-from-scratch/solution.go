package minheapfromscratch

// MinHeap is a binary min-heap stored in a slice.
type MinHeap struct{ a []int }

// Build creates a heap from values in O(n) by sifting down every non-leaf.
func Build(values []int) *MinHeap {
	h := &MinHeap{a: append([]int{}, values...)}
	for i := len(h.a)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

// Len returns the number of elements.
func (h *MinHeap) Len() int { return len(h.a) }

// Peek returns the minimum. O(1).
func (h *MinHeap) Peek() (int, bool) {
	if len(h.a) == 0 {
		return 0, false
	}
	return h.a[0], true
}

// Push inserts x. O(log n).
func (h *MinHeap) Push(x int) {
	h.a = append(h.a, x)
	h.siftUp(len(h.a) - 1)
}

// Pop removes and returns the minimum. O(log n).
func (h *MinHeap) Pop() (int, bool) {
	if len(h.a) == 0 {
		return 0, false
	}
	top := h.a[0]
	last := len(h.a) - 1
	h.a[0] = h.a[last]
	h.a = h.a[:last]
	if len(h.a) > 0 {
		h.siftDown(0)
	}
	return top, true
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.a[p] <= h.a[i] {
			return
		}
		h.a[p], h.a[i] = h.a[i], h.a[p]
		i = p
	}
}

func (h *MinHeap) siftDown(i int) {
	n := len(h.a)
	for {
		smallest := i
		l, r := 2*i+1, 2*i+2
		if l < n && h.a[l] < h.a[smallest] {
			smallest = l
		}
		if r < n && h.a[r] < h.a[smallest] {
			smallest = r
		}
		if smallest == i {
			return
		}
		h.a[i], h.a[smallest] = h.a[smallest], h.a[i]
		i = smallest
	}
}

// HeapSort sorts values ascending using the heap. O(n log n).
func HeapSort(values []int) []int {
	h := Build(values)
	out := make([]int, 0, len(values))
	for h.Len() > 0 {
		x, _ := h.Pop()
		out = append(out, x)
	}
	return out
}
