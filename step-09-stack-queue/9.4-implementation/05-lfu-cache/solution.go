package lfucache

type node struct {
	key, val, freq int
	prev, next     *node
}

// dlist is a doubly linked list with dummy head/tail. Front = most recent.
type dlist struct {
	head, tail *node
	size       int
}

func newList() *dlist {
	h, t := &node{}, &node{}
	h.next, t.prev = t, h
	return &dlist{head: h, tail: t}
}

func (l *dlist) addFront(n *node) {
	n.prev, n.next = l.head, l.head.next
	l.head.next.prev = n
	l.head.next = n
	l.size++
}

func (l *dlist) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--
}

// LFUCache evicts the least frequently used key (LRU among ties). O(1) ops.
type LFUCache struct {
	cap, minFreq int
	items        map[int]*node
	lists        map[int]*dlist
}

// Constructor creates a cache with the given capacity.
func Constructor(capacity int) LFUCache {
	return LFUCache{cap: capacity, items: map[int]*node{}, lists: map[int]*dlist{}}
}

func (c *LFUCache) list(f int) *dlist {
	if c.lists[f] == nil {
		c.lists[f] = newList()
	}
	return c.lists[f]
}

// touch moves n from frequency f to f+1.
func (c *LFUCache) touch(n *node) {
	l := c.lists[n.freq]
	l.remove(n)
	if l.size == 0 && n.freq == c.minFreq {
		c.minFreq++
	}
	n.freq++
	c.list(n.freq).addFront(n)
}

// Get returns the value for key, or -1.
func (c *LFUCache) Get(key int) int {
	n, ok := c.items[key]
	if !ok {
		return -1
	}
	c.touch(n)
	return n.val
}

// Put inserts or updates key.
func (c *LFUCache) Put(key, value int) {
	if c.cap == 0 {
		return
	}
	if n, ok := c.items[key]; ok {
		n.val = value
		c.touch(n)
		return
	}
	if len(c.items) == c.cap {
		l := c.lists[c.minFreq]
		victim := l.tail.prev // least recently used among the least frequent
		l.remove(victim)
		delete(c.items, victim.key)
	}
	n := &node{key: key, val: value, freq: 1}
	c.items[key] = n
	c.list(1).addFront(n)
	c.minFreq = 1
}
