package lrucache

import "testing"

func TestLRU(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Error("get 1")
	}
	c.Put(3, 3) // evicts 2
	if c.Get(2) != -1 {
		t.Error("2 should be evicted")
	}
	c.Put(4, 4) // evicts 1
	if c.Get(1) != -1 || c.Get(3) != 3 || c.Get(4) != 4 {
		t.Error("after put 4")
	}
	c.Put(3, 30) // update existing
	if c.Get(3) != 30 {
		t.Error("update")
	}
}
