package sortstack

// Stack is a slice where the last element is the top.
type Stack []int

func (s *Stack) push(x int) { *s = append(*s, x) }
func (s *Stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s Stack) top() int { return s[len(s)-1] }

// Sort sorts the stack (largest on top) using only recursion. Time O(n^2).
func Sort(s *Stack) {
	if len(*s) == 0 {
		return
	}
	x := s.pop()
	Sort(s)
	insert(s, x)
}

// insert puts x into an already sorted stack.
func insert(s *Stack, x int) {
	if len(*s) == 0 || s.top() <= x {
		s.push(x)
		return
	}
	t := s.pop()
	insert(s, x)
	s.push(t)
}
