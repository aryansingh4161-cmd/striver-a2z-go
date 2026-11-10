package stackusingarray

import "testing"

func TestArrayStack(t *testing.T) {
	s := NewArrayStack(2)
	if _, ok := s.Pop(); ok {
		t.Error("pop on empty")
	}
	s.Push(1)
	s.Push(2)
	if s.Push(3) {
		t.Error("overflow not detected")
	}
	if x, _ := s.Top(); x != 2 || s.Size() != 2 {
		t.Error("top/size")
	}
	if x, _ := s.Pop(); x != 2 {
		t.Error("pop")
	}
	if x, _ := s.Pop(); x != 1 || s.Size() != 0 {
		t.Error("pop 2")
	}
}

func TestSliceStack(t *testing.T) {
	var s SliceStack
	for i := 1; i <= 5; i++ {
		s.Push(i)
	}
	for want := 5; want >= 1; want-- {
		if x, ok := s.Pop(); !ok || x != want {
			t.Fatalf("pop got %d want %d", x, want)
		}
	}
	if _, ok := s.Top(); ok {
		t.Error("top on empty")
	}
}
