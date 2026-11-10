package queueusingstack

import "testing"

func TestMyQueue(t *testing.T) {
	var q MyQueue
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 || q.Pop() != 1 {
		t.Error("front should be 1")
	}
	q.Push(3)
	if q.Pop() != 2 || q.Pop() != 3 || !q.Empty() {
		t.Error("order 2, 3")
	}
}
