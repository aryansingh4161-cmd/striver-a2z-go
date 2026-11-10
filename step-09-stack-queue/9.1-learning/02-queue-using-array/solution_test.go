package queueusingarray

import "testing"

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(3)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.Push(4) {
		t.Error("should be full")
	}
	if x, _ := q.Pop(); x != 1 {
		t.Error("pop 1")
	}
	q.Push(4) // wraps around
	want := []int{2, 3, 4}
	for _, w := range want {
		if x, ok := q.Pop(); !ok || x != w {
			t.Fatalf("got %d want %d", x, w)
		}
	}
	if _, ok := q.Pop(); ok || q.Size() != 0 {
		t.Error("should be empty")
	}
}
