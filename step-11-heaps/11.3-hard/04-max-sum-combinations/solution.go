package maxsumcombinations

import (
	"container/heap"
	"sort"
)

type item struct{ sum, i, j int }

type maxHeap []item

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].sum > h[j].sum }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// TopKSums returns the k largest a[i]+b[j] sums in decreasing order.
// Time O(n log n + k log k).
func TopKSums(a, b []int, k int) []int {
	x := append([]int{}, a...)
	y := append([]int{}, b...)
	sort.Sort(sort.Reverse(sort.IntSlice(x)))
	sort.Sort(sort.Reverse(sort.IntSlice(y)))
	res := []int{}
	if len(x) == 0 || len(y) == 0 {
		return res
	}
	h := &maxHeap{{x[0] + y[0], 0, 0}}
	seen := map[[2]int]bool{{0, 0}: true}
	for len(res) < k && h.Len() > 0 {
		it := heap.Pop(h).(item)
		res = append(res, it.sum)
		for _, nb := range [][2]int{{it.i + 1, it.j}, {it.i, it.j + 1}} {
			if nb[0] < len(x) && nb[1] < len(y) && !seen[nb] {
				seen[nb] = true
				heap.Push(h, item{x[nb[0]] + y[nb[1]], nb[0], nb[1]})
			}
		}
	}
	return res
}
