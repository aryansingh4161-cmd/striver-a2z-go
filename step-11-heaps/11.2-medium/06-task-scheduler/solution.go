package taskscheduler

import "container/heap"

// LeastIntervalFormula uses the block formula. Time O(n), space O(26).
func LeastIntervalFormula(tasks []byte, n int) int {
	var freq [26]int
	maxFreq := 0
	for _, t := range tasks {
		freq[t-'A']++
		maxFreq = max(maxFreq, freq[t-'A'])
	}
	countMax := 0
	for _, f := range freq {
		if f == maxFreq {
			countMax++
		}
	}
	return max(len(tasks), (maxFreq-1)*(n+1)+countMax)
}

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

// LeastIntervalHeap simulates rounds of n+1 slots with a max-heap of counts.
func LeastIntervalHeap(tasks []byte, n int) int {
	var freq [26]int
	for _, t := range tasks {
		freq[t-'A']++
	}
	h := &IntHeap{} // stores negative counts → acts as a max-heap
	for _, f := range freq {
		if f > 0 {
			heap.Push(h, -f)
		}
	}
	time := 0
	for h.Len() > 0 {
		var back []int
		ran := 0
		for i := 0; i <= n && h.Len() > 0; i++ {
			c := -heap.Pop(h).(int) - 1
			ran++
			if c > 0 {
				back = append(back, c)
			}
		}
		for _, c := range back {
			heap.Push(h, -c)
		}
		if h.Len() == 0 {
			time += ran // last round: no idle at the end
		} else {
			time += n + 1
		}
	}
	return time
}
