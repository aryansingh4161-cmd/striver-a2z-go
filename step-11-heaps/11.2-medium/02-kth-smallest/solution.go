package kthsmallest

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

// KthSmallest keeps the k smallest in a max-heap (stored as negatives in a
// min-heap). Time O(n log k), space O(k).
func KthSmallest(a []int, k int) int {
	h := &IntHeap{}
	for _, x := range a {
		heap.Push(h, -x)
		if h.Len() > k {
			heap.Pop(h) // removes the most negative = the largest value
		}
	}
	return -(*h)[0]
}
