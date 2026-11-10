package queueusinglinkedlist

import "testing"

func TestLinkedQueue(t *testing.T) {
	var q LinkedQueue
	q.Push(1)
	q.Push(2)
	if x, _ := q.Pop(); x != 1 {
		t.Error("pop 1")
	}
	if x, _ := q.Pop(); x != 2 {
		t.Error("pop 2")
	}
	q.Push(3) // works again after becoming empty
	if x, _ := q.Front(); x != 3 || q.Size() != 1 {
		t.Error("front 3")
	}
	q.Pop()
	if _, ok := q.Pop(); ok {
		t.Error("underflow")
	}
}
