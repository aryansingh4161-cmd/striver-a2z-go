# Step 6: Linked List

> Status: theory ready, problems coming ⏳

## What it is
A chain of **nodes**. Each node holds a value and a pointer to the next node. Unlike a slice, nodes are scattered in memory.

```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

| Operation | Slice | Linked list |
|---|---|---|
| Access i-th element | O(1) | O(n) |
| Insert/delete at the front | O(n) | **O(1)** |
| Insert/delete after a known node | O(n) | **O(1)** |
| Extra memory | none | one pointer per node |

**Doubly linked list (DLL):** each node also has `Prev`. Lets you walk backwards and delete a node in O(1) if you have a pointer to it (used in LRU cache).

## Basic walks
```go
// traverse
for cur := head; cur != nil; cur = cur.Next {
    fmt.Println(cur.Val)
}

// insert at head
head = &ListNode{Val: 5, Next: head}
```

## The tricks that solve most problems
1. **Dummy node:** create `dummy := &ListNode{Next: head}` so deleting/inserting at the head isn't a special case. Return `dummy.Next`.
2. **Slow and fast pointers (tortoise and hare):** slow moves 1, fast moves 2.
   - fast reaches the end → slow is at the **middle**.
   - fast meets slow → there's a **cycle**. Reset one to head and move both by 1; they meet at the **cycle start**.
3. **Reverse a list (iterative):**
   ```go
   var prev *ListNode
   for cur := head; cur != nil; {
       next := cur.Next
       cur.Next = prev
       prev, cur = cur, next
   }
   return prev
   ```
4. **Two pointers with a gap:** to remove the n-th node from the end, move one pointer n steps ahead first.
5. Draw boxes and arrows on paper. Most linked-list bugs come from losing a pointer before you've saved it.

## Problems

### 6.1 Singly linked list basics
| # | Problem | Status |
|---|---|---|
| 1 | Convert slice to linked list | ⏳ |
| 2 | Insert (head, tail, k-th position, before value) | ⏳ |
| 3 | Delete (head, tail, k-th, by value) | ⏳ |
| 4 | Length of a linked list | ⏳ |
| 5 | Search an element | ⏳ |

### 6.2 Doubly linked list basics
| # | Problem | Status |
|---|---|---|
| 1 | Convert slice to DLL | ⏳ |
| 2 | Insert in DLL | ⏳ |
| 3 | Delete in DLL | ⏳ |
| 4 | Reverse a DLL | ⏳ |

### 6.3 Medium (singly)
| # | Problem | Status |
|---|---|---|
| 1 | Middle of the linked list | ⏳ |
| 2 | Reverse a linked list (iterative + recursive) | ⏳ |
| 3 | Detect a loop | ⏳ |
| 4 | Starting point of the loop | ⏳ |
| 5 | Length of the loop | ⏳ |
| 6 | Check if palindrome | ⏳ |
| 7 | Odd-even linked list | ⏳ |
| 8 | Remove N-th node from the end | ⏳ |
| 9 | Delete the middle node | ⏳ |
| 10 | Sort a linked list | ⏳ |
| 11 | Sort a list of 0s, 1s, 2s | ⏳ |
| 12 | Intersection point of two lists | ⏳ |
| 13 | Add 1 to a number as a linked list | ⏳ |
| 14 | Add two numbers | ⏳ |

### 6.4 Medium (doubly)
| # | Problem | Status |
|---|---|---|
| 1 | Delete all occurrences of a key | ⏳ |
| 2 | Find pairs with a given sum in sorted DLL | ⏳ |
| 3 | Remove duplicates from sorted DLL | ⏳ |

### 6.5 Hard
| # | Problem | Status |
|---|---|---|
| 1 | Reverse nodes in K groups | ⏳ |
| 2 | Rotate a linked list | ⏳ |
| 3 | Flatten a linked list | ⏳ |
| 4 | Clone a linked list with random pointer | ⏳ |
