package lrucache

type node struct {
	key, val   int
	prev, next *node
}

// LRUCache evicts the least recently used key. All operations O(1).
type LRUCache struct {
	cap        int
	items      map[int]*node
	head, tail *node // dummies: head.next = most recent, tail.prev = least recent
}

// Constructor creates a cache with the given capacity.
func Constructor(capacity int) LRUCache {
	h, t := &node{}, &node{}
	h.next, t.prev = t, h
	return LRUCache{cap: capacity, items: map[int]*node{}, head: h, tail: t}
}

func (c *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) addFront(n *node) {
	n.prev, n.next = c.head, c.head.next
	c.head.next.prev = n
	c.head.next = n
}

// Get returns the value for key (or -1) and marks it most recently used.
func (c *LRUCache) Get(key int) int {
	n, ok := c.items[key]
	if !ok {
		return -1
	}
	c.remove(n)
	c.addFront(n)
	return n.val
}

// Put inserts or updates key, evicting the least recently used if needed.
func (c *LRUCache) Put(key, value int) {
	if n, ok := c.items[key]; ok {
		n.val = value
		c.remove(n)
		c.addFront(n)
		return
	}
	if len(c.items) == c.cap {
		lru := c.tail.prev
		c.remove(lru)
		delete(c.items, lru.key)
	}
	n := &node{key: key, val: value}
	c.items[key] = n
	c.addFront(n)
}
