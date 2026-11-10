package lfucache

import "testing"

func TestLFU(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 { // freq(1) = 2
		t.Error("get 1")
	}
	c.Put(3, 3) // evicts 2 (freq 1)
	if c.Get(2) != -1 {
		t.Error("2 should be evicted")
	}
	if c.Get(3) != 3 { // freq(3) = 2
		t.Error("get 3")
	}
	c.Put(4, 4) // tie freq 2 between 1 and 3 → evict LRU = 1
	if c.Get(1) != -1 || c.Get(3) != 3 || c.Get(4) != 4 {
		t.Error("after put 4")
	}
	z := Constructor(0)
	z.Put(1, 1)
	if z.Get(1) != -1 {
		t.Error("capacity 0")
	}
}
