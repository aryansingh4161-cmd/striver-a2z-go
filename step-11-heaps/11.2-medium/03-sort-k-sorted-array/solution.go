package sortksortedarray

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

// SortKSorted sorts an array where each element is at most k away from its
// final position. Time O(n log k), space O(k).
func SortKSorted(a []int, k int) []int {
	h := &IntHeap{}
	out := make([]int, 0, len(a))
	for _, x := range a {
		heap.Push(h, x)
		if h.Len() > k {
			out = append(out, heap.Pop(h).(int))
		}
	}
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(int))
	}
	return out
}
