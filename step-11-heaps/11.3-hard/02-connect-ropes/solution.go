package connectropes

import "container/heap"

// IntHeap is a min-heap of ints for container/heap.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// MinCost joins ropes, always taking the two shortest. Time O(n log n).
func MinCost(ropes []int) int {
	h := &IntHeap{}
	*h = append(*h, ropes...)
	heap.Init(h)
	cost := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		cost += a + b
		heap.Push(h, a+b)
	}
	return cost
}
