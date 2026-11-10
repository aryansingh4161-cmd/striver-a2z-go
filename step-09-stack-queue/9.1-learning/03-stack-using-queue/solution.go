package stackusingqueue

// queue is a minimal FIFO queue (only push back / pop front are used).
type queue []int

func (q *queue) push(x int) { *q = append(*q, x) }
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

// MyStack is a LIFO stack built from one queue.
type MyStack struct{ q queue }

// Push adds x on top. O(n): rotates older elements behind x.
func (s *MyStack) Push(x int) {
	s.q.push(x)
	for i := 0; i < len(s.q)-1; i++ {
		s.q.push(s.q.pop())
	}
}

// Pop removes and returns the top. O(1).
func (s *MyStack) Pop() int { return s.q.pop() }

// Top returns the top element. O(1).
func (s *MyStack) Top() int { return s.q[0] }

// Empty reports whether the stack is empty.
func (s *MyStack) Empty() bool { return len(s.q) == 0 }
