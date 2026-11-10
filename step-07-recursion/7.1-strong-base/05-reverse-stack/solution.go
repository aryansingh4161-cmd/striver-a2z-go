package reversestack

// Stack is a slice where the last element is the top.
type Stack []int

func (s *Stack) push(x int) { *s = append(*s, x) }
func (s *Stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

// Reverse reverses the stack using only recursion. Time O(n^2).
func Reverse(s *Stack) {
	if len(*s) == 0 {
		return
	}
	x := s.pop()
	Reverse(s)
	insertBottom(s, x)
}

func insertBottom(s *Stack, x int) {
	if len(*s) == 0 {
		s.push(x)
		return
	}
	t := s.pop()
	insertBottom(s, x)
	s.push(t)
}
