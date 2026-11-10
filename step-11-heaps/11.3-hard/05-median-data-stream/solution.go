package mediandatastream

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

// MedianFinder keeps the lower half in a max-heap and the upper half in a min-heap.
type MedianFinder struct {
	low  *IntHeap // max-heap via negated values
	high *IntHeap // min-heap
}

// Constructor creates an empty MedianFinder.
func Constructor() MedianFinder {
	return MedianFinder{low: &IntHeap{}, high: &IntHeap{}}
}

// AddNum adds x. O(log n).
func (m *MedianFinder) AddNum(x int) {
	heap.Push(m.low, -x)
	heap.Push(m.high, -heap.Pop(m.low).(int))
	if m.high.Len() > m.low.Len() {
		heap.Push(m.low, -heap.Pop(m.high).(int))
	}
}

// FindMedian returns the current median. O(1).
func (m *MedianFinder) FindMedian() float64 {
	if m.low.Len() > m.high.Len() {
		return float64(-(*m.low)[0])
	}
	return float64(-(*m.low)[0]+(*m.high)[0]) / 2
}
