package topkfrequent

import (
	"container/heap"
	"sort"
)

type pair struct{ freq, val int }

type minHeap []pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// TopKHeap keeps the k most frequent in a min-heap. O(n log k). Result sorted.
func TopKHeap(a []int, k int) []int {
	freq := map[int]int{}
	for _, x := range a {
		freq[x]++
	}
	h := &minHeap{}
	for v, f := range freq {
		heap.Push(h, pair{f, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := []int{}
	for _, p := range *h {
		res = append(res, p.val)
	}
	sort.Ints(res)
	return res
}

// TopKBucket uses buckets indexed by frequency. O(n). Result sorted.
func TopKBucket(a []int, k int) []int {
	freq := map[int]int{}
	for _, x := range a {
		freq[x]++
	}
	buckets := make([][]int, len(a)+1)
	for v, f := range freq {
		buckets[f] = append(buckets[f], v)
	}
	res := []int{}
	for f := len(a); f >= 1 && len(res) < k; f-- {
		res = append(res, buckets[f]...)
	}
	res = res[:min(k, len(res))]
	sort.Ints(res)
	return res
}
