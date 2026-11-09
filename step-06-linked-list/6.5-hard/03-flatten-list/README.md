# Flatten a Linked List

**Problem:** Each node has `Next` (to the next head on the top level) and `Child` (down a vertical sorted list). Every vertical list is sorted. Flatten everything into **one sorted list** using only `Child` pointers.

```
5 → 10 → 19 → 28
↓    ↓    ↓    ↓
7   20   22   35
↓         ↓    ↓
8        50   40
↓              ↓
30            45
```
→ `5 7 8 10 19 20 22 28 30 35 40 45 50`

## Approach 1: Collect, sort, rebuild
Put every value in a slice, sort, build a new child-list. **Time O(N log N), Space O(N)** (N = total nodes).

## Approach 2: Merge from the right (optimal)
This is "merge many sorted lists". Recursively flatten everything to the right of `head`, then **merge** `head`'s vertical list with it (same merge as merge sort, but on `Child` pointers).
```
flatten(head):
    if head == nil or head.Next == nil: return head
    right = flatten(head.Next)
    head.Next = nil
    return merge(head, right)   // merge via Child
```
- **Time:** O(N × k) where k = number of vertical lists (each merge walks the growing result)
- **Space:** O(k) recursion stack

(A min-heap version is O(N log k). Heaps come in Step 11.)
