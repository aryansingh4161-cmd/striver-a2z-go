package queueusingarray

// CircularQueue is a fixed-capacity FIFO queue backed by an array.
type CircularQueue struct {
	arr                []int
	front, rear, count int
}

// NewCircularQueue creates a queue of the given capacity.
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{arr: make([]int, capacity), rear: -1}
}

// Push adds x at the back. Returns false if full. O(1).
func (q *CircularQueue) Push(x int) bool {
	if q.count == len(q.arr) {
		return false
	}
	q.rear = (q.rear + 1) % len(q.arr)
	q.arr[q.rear] = x
	q.count++
	return true
}

// Pop removes and returns the front. ok=false if empty. O(1).
func (q *CircularQueue) Pop() (int, bool) {
	if q.count == 0 {
		return 0, false
	}
	x := q.arr[q.front]
	q.front = (q.front + 1) % len(q.arr)
	q.count--
	return x, true
}

// Front returns the front without removing it. O(1).
func (q *CircularQueue) Front() (int, bool) {
	if q.count == 0 {
		return 0, false
	}
	return q.arr[q.front], true
}

// Size returns the number of elements. O(1).
func (q *CircularQueue) Size() int { return q.count }
