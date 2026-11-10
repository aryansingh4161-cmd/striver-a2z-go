package stackusingqueue

import "testing"

func TestMyStack(t *testing.T) {
	var s MyStack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Top() != 3 || s.Pop() != 3 || s.Pop() != 2 {
		t.Error("LIFO order broken")
	}
	s.Push(9)
	if s.Pop() != 9 || s.Pop() != 1 || !s.Empty() {
		t.Error("after push 9")
	}
}
