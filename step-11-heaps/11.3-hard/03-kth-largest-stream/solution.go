package kthlargeststream

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

// KthLargest tracks the k-th largest value of a stream.
type KthLargest struct {
	k int
	h *IntHeap
}

// Constructor builds the tracker from the initial numbers.
func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k, h: &IntHeap{}}
	for _, x := range nums {
		kl.Add(x)
	}
	return kl
}

// Add inserts x and returns the current k-th largest. O(log k).
func (kl *KthLargest) Add(x int) int {
	heap.Push(kl.h, x)
	if kl.h.Len() > kl.k {
		heap.Pop(kl.h)
	}
	return (*kl.h)[0]
}
