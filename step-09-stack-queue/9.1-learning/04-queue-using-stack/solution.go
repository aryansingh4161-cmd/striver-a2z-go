package queueusingstack

// MyQueue is a FIFO queue built from two stacks (slices).
type MyQueue struct{ in, out []int }

// Push adds x to the back. O(1).
func (q *MyQueue) Push(x int) { q.in = append(q.in, x) }

// move pours in into out when out is empty.
func (q *MyQueue) move() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			top := q.in[len(q.in)-1]
			q.in = q.in[:len(q.in)-1]
			q.out = append(q.out, top)
		}
	}
}

// Pop removes and returns the front. Amortised O(1).
func (q *MyQueue) Pop() int {
	q.move()
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

// Peek returns the front. Amortised O(1).
func (q *MyQueue) Peek() int {
	q.move()
	return q.out[len(q.out)-1]
}

// Empty reports whether the queue is empty.
func (q *MyQueue) Empty() bool { return len(q.in) == 0 && len(q.out) == 0 }
