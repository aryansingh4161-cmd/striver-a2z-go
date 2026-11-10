package minstack

// MinStack stores (value, minimum so far) pairs. All operations O(1).
type MinStack struct{ st [][2]int }

// Push adds x.
func (s *MinStack) Push(x int) {
	m := x
	if len(s.st) > 0 && s.st[len(s.st)-1][1] < m {
		m = s.st[len(s.st)-1][1]
	}
	s.st = append(s.st, [2]int{x, m})
}

// Pop removes the top element.
func (s *MinStack) Pop() { s.st = s.st[:len(s.st)-1] }

// Top returns the top value.
func (s *MinStack) Top() int { return s.st[len(s.st)-1][0] }

// GetMin returns the current minimum.
func (s *MinStack) GetMin() int { return s.st[len(s.st)-1][1] }

// MinStackEncoded stores one value per entry using the 2x - min trick.
type MinStackEncoded struct {
	st   []int
	mini int
}

// Push adds x.
func (s *MinStackEncoded) Push(x int) {
	switch {
	case len(s.st) == 0:
		s.st = append(s.st, x)
		s.mini = x
	case x < s.mini:
		s.st = append(s.st, 2*x-s.mini) // encoded, always < x
		s.mini = x
	default:
		s.st = append(s.st, x)
	}
}

// Pop removes the top element.
func (s *MinStackEncoded) Pop() {
	v := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	if v < s.mini { // encoded: restore the previous minimum
		s.mini = 2*s.mini - v
	}
}

// Top returns the top value.
func (s *MinStackEncoded) Top() int {
	v := s.st[len(s.st)-1]
	if v < s.mini {
		return s.mini
	}
	return v
}

// GetMin returns the current minimum.
func (s *MinStackEncoded) GetMin() int { return s.mini }
