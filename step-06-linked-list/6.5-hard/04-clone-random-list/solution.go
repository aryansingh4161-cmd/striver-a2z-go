package clonerandomlist

// Node has an extra Random pointer.
type Node struct {
	Val          int
	Next, Random *Node
}

// CloneMap uses a map from old node to new node. Time O(n), space O(n).
func CloneMap(head *Node) *Node {
	m := map[*Node]*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}
	return m[head]
}

// Clone interleaves copies with originals. Time O(n), space O(1) extra.
func Clone(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. insert copies
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	// 2. set random pointers
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	// 3. separate
	dummy := &Node{}
	tail := dummy
	for cur := head; cur != nil; cur = cur.Next {
		copyNode := cur.Next
		cur.Next = copyNode.Next // restore original
		tail.Next = copyNode
		tail = copyNode
	}
	return dummy.Next
}
