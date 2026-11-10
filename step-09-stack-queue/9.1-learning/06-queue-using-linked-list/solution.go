package queueusinglinkedlist

type node struct {
	val  int
	next *node
}

// LinkedQueue is a FIFO queue with head (front) and tail (back) pointers.
type LinkedQueue struct {
	head, tail *node
	size       int
}

// Push adds x at the back. O(1).
func (q *LinkedQueue) Push(x int) {
	n := &node{val: x}
	if q.tail == nil {
		q.head, q.tail = n, n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.size++
}

// Pop removes and returns the front. O(1).
func (q *LinkedQueue) Pop() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	x := q.head.val
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return x, true
}

// Front returns the front. O(1).
func (q *LinkedQueue) Front() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	return q.head.val, true
}

// Size returns the number of elements.
func (q *LinkedQueue) Size() int { return q.size }
