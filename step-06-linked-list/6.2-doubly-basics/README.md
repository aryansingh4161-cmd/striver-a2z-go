# 6.2 Doubly Linked List Basics

Each node has `Prev` and `Next`. Code in [`basics.go`](basics.go).

```go
type DNode struct {
    Val        int
    Prev, Next *DNode
}
```

**Why use it?** You can walk backwards, and if you already hold a pointer to a node you can delete it in **O(1)** without searching for its previous node. That's what makes an LRU cache O(1) (Step 9).

**Cost:** one more pointer per node, and every insert/delete must fix **both** directions.

## Insert / delete rules
When you link node `n` between `a` and `b`, set four pointers:
`n.Prev = a`, `n.Next = b`, `a.Next = n`, `b.Prev = n` (skip the ones where `a` or `b` is nil).

When you delete `n`:
`n.Prev.Next = n.Next` and `n.Next.Prev = n.Prev` (again skipping nil neighbours). If `n` was the head, the new head is `n.Next`.

## Reverse a DLL
For every node, **swap** its `Prev` and `Next`. The last node visited becomes the new head.
```go
var last *DNode
for cur := head; cur != nil; cur = cur.Prev { // after the swap, "Prev" is the old Next
    cur.Prev, cur.Next = cur.Next, cur.Prev
    last = cur
}
return last
```
- **Time:** O(n), **Space:** O(1)

(The brute force pushes all values onto a stack and writes them back: O(n) space.)
