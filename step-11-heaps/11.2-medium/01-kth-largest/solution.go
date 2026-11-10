package kthlargest

import (
	"container/heap"
	"math/rand"
)

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

// KthLargestHeap keeps a min-heap of the k largest. Time O(n log k), space O(k).
func KthLargestHeap(a []int, k int) int {
	h := &IntHeap{}
	for _, x := range a {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// KthLargestQuickselect uses partitioning with a random pivot. Average O(n).
func KthLargestQuickselect(nums []int, k int) int {
	a := append([]int{}, nums...)
	target := len(a) - k // index in ascending order
	lo, hi := 0, len(a)-1
	for {
		p := partition(a, lo, hi)
		switch {
		case p == target:
			return a[p]
		case p < target:
			lo = p + 1
		default:
			hi = p - 1
		}
	}
}

// partition (Lomuto) with a random pivot; returns the pivot's final index.
func partition(a []int, lo, hi int) int {
	r := lo + rand.Intn(hi-lo+1)
	a[r], a[hi] = a[hi], a[r]
	pivot, i := a[hi], lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
