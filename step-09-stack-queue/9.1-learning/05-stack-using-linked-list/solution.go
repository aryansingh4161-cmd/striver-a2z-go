package stackusinglinkedlist

type node struct {
	val  int
	next *node
}

// LinkedStack is a stack using a singly linked list (top = head).
type LinkedStack struct {
	top  *node
	size int
}

// Push adds x on top. O(1).
func (s *LinkedStack) Push(x int) {
	s.top = &node{val: x, next: s.top}
	s.size++
}

// Pop removes and returns the top. O(1).
func (s *LinkedStack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	x := s.top.val
	s.top = s.top.next
	s.size--
	return x, true
}

// Top returns the top. O(1).
func (s *LinkedStack) Top() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.val, true
}

// Size returns the number of elements.
func (s *LinkedStack) Size() int { return s.size }
