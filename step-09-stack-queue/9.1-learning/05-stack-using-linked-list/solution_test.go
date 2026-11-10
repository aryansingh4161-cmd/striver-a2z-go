package stackusinglinkedlist

import "testing"

func TestLinkedStack(t *testing.T) {
	var s LinkedStack
	s.Push(1)
	s.Push(2)
	if x, _ := s.Top(); x != 2 || s.Size() != 2 {
		t.Error("top/size")
	}
	if x, _ := s.Pop(); x != 2 {
		t.Error("pop 2")
	}
	if x, _ := s.Pop(); x != 1 {
		t.Error("pop 1")
	}
	if _, ok := s.Pop(); ok {
		t.Error("underflow")
	}
}
