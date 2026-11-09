# 6.1 Singly Linked List Basics

All the basic operations live in [`basics.go`](basics.go). The node type comes from the shared [`list`](../list/) package:

```go
type Node struct {
    Val  int
    Next *Node
}
```

## Slice → linked list
Keep a `tail` pointer; attach each new node to `tail.Next`. A **dummy** head node avoids a special case for the first element.
**O(n)**.

## Length and search
Walk with `for cur := head; cur != nil; cur = cur.Next`. Count, or stop when you find the value. **O(n)**.

## Insertion
| Where | How | Time |
|---|---|---|
| head | `return &Node{Val: v, Next: head}` | O(1) |
| tail | walk to the last node, attach | O(n) |
| position k (1-based) | walk to node k−1, then `new.Next = prev.Next; prev.Next = new` | O(k) |
| before value x | walk until `cur.Next.Val == x`, insert after `cur` | O(n) |

**Order matters:** set `new.Next` **before** changing `prev.Next`, or you lose the rest of the list.

## Deletion
| What | How | Time |
|---|---|---|
| head | `return head.Next` | O(1) |
| tail | walk to the second-last node, set its `Next = nil` | O(n) |
| k-th node | walk to node k−1, `prev.Next = prev.Next.Next` | O(k) |
| first node with value x | same, stop at the node before x | O(n) |

In Go there's no `delete`/`free`: once nothing points to a node, the garbage collector removes it.

## Edge cases to always test
- empty list (`head == nil`)
- one node
- operating on the head or the tail
- position larger than the length
