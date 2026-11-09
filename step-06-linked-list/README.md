# Step 6: Linked List

> Status: ✅ theory and all problems done

Shared helpers live in [`list/`](list/): node types plus `FromSlice` / `ToSlice` (and doubly-linked versions) so tests can build lists in one line.

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Convert slice to linked list](6.1-singly-basics/) | dummy + tail pointer (list.FromSlice) |
| 2 | [Insert: head, tail, k-th, before value](6.1-singly-basics/) | set new.Next before prev.Next |
| 3 | [Delete: head, tail, k-th, by value](6.1-singly-basics/) | stop at the node before |
| 4 | [Length](6.1-singly-basics/) | walk and count |
| 5 | [Search](6.1-singly-basics/) | walk and compare |

### 6.2 Doubly linked list basics
| # | Problem | Key idea |
|---|---|---|
| 1 | [Convert slice to DLL](6.2-doubly-basics/) | link Prev and Next |
| 2 | [Insert in DLL](6.2-doubly-basics/) | fix four pointers |
| 3 | [Delete in DLL](6.2-doubly-basics/) | unlink both sides |
| 4 | [Reverse a DLL](6.2-doubly-basics/) | swap Prev and Next |

### 6.3 Medium (singly)
| # | Problem | Key idea |
|---|---|---|
| 1 | [Middle of the list](6.3-medium/01-middle-of-list/) | slow/fast |
| 2 | [Reverse a list](6.3-medium/02-reverse-list/) | prev, cur, next / recursion |
| 3 | [Detect a loop](6.3-medium/03-detect-loop/) | Floyd: slow meets fast |
| 4 | [Start of the loop](6.3-medium/04-start-of-loop/) | reset slow to head |
| 5 | [Length of the loop](6.3-medium/05-length-of-loop/) | count one lap |
| 6 | [Palindrome list](6.3-medium/06-palindrome-list/) | reverse second half |
| 7 | [Odd-even list](6.3-medium/07-odd-even-list/) | two chains |
| 8 | [Remove N-th from end](6.3-medium/08-remove-nth-from-end/) | gap of n |
| 9 | [Delete the middle](6.3-medium/09-delete-middle/) | fast starts 2 ahead |
| 10 | [Sort a list](6.3-medium/10-sort-list/) | merge sort on the list |
| 11 | [Sort 0s, 1s, 2s](6.3-medium/11-sort-0s-1s-2s/) | three dummy chains |
| 12 | [Intersection point](6.3-medium/12-intersection-point/) | switch heads |
| 13 | [Add 1 to a number](6.3-medium/13-add-one/) | reverse or recursive carry |
| 14 | [Add two numbers](6.3-medium/14-add-two-numbers/) | digit-by-digit with carry |

### 6.4 Medium (doubly)
| # | Problem | Key idea |
|---|---|---|
| 1 | [Delete all occurrences of a key](6.4-doubly-medium/01-delete-all-occurrences/) | save next before unlinking |
| 2 | [Pairs with given sum (sorted DLL)](6.4-doubly-medium/02-pairs-with-sum/) | two pointers from both ends |
| 3 | [Remove duplicates (sorted DLL)](6.4-doubly-medium/03-remove-duplicates-sorted-dll/) | skip equal run, relink |

### 6.5 Hard
| # | Problem | Key idea |
|---|---|---|
| 1 | [Reverse nodes in K groups](6.5-hard/01-reverse-k-group/) | reverse group, reconnect |
| 2 | [Rotate a list](6.5-hard/02-rotate-list/) | make a ring, cut at n−k |
| 3 | [Flatten a list](6.5-hard/03-flatten-list/) | merge vertical lists from the right |
| 4 | [Clone list with random pointer](6.5-hard/04-clone-random-list/) | interleave copies |
