package stackusingarray

// ArrayStack is a fixed-capacity stack backed by an array.
type ArrayStack struct {
	arr []int
	top int // index of the top element, -1 when empty
}

// NewArrayStack creates a stack that can hold up to capacity items.
func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{arr: make([]int, capacity), top: -1}
}

// Push adds x. Returns false on overflow. O(1).
func (s *ArrayStack) Push(x int) bool {
	if s.top == len(s.arr)-1 {
		return false
	}
	s.top++
	s.arr[s.top] = x
	return true
}

// Pop removes and returns the top. ok=false on underflow. O(1).
func (s *ArrayStack) Pop() (int, bool) {
	if s.top == -1 {
		return 0, false
	}
	x := s.arr[s.top]
	s.top--
	return x, true
}

// Top returns the top without removing it. O(1).
func (s *ArrayStack) Top() (int, bool) {
	if s.top == -1 {
		return 0, false
	}
	return s.arr[s.top], true
}

// Size returns the number of elements. O(1).
func (s *ArrayStack) Size() int { return s.top + 1 }

// SliceStack is the idiomatic Go stack: a growing slice.
type SliceStack []int

func (s *SliceStack) Push(x int) { *s = append(*s, x) }

func (s *SliceStack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x, true
}

func (s SliceStack) Top() (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	return s[len(s)-1], true
}

func (s SliceStack) Size() int { return len(s) }
