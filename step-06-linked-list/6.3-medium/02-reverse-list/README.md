# Reverse a Linked List

**Problem (LeetCode 206):** Reverse the list and return the new head.

`1 → 2 → 3 → 4` → `4 → 3 → 2 → 1`

## Approach 1: Stack of values (brute)
Push all values, then overwrite the nodes while popping. **Time O(2n), Space O(n)**.

## Approach 2: Iterative pointer reversal (optimal)
Turn every arrow around. Keep three pointers: `prev`, `cur`, `next`.
```
prev = nil
while cur != nil:
    next = cur.Next      // 1. save the rest
    cur.Next = prev      // 2. flip the arrow
    prev = cur           // 3. move prev forward
    cur = next           // 4. move cur forward
return prev
```
- **Time:** O(n), **Space:** O(1)

## Dry run `1 → 2 → 3`
| step | prev | cur | list state |
|---|---|---|---|
| start | nil | 1 | 1→2→3 |
| 1 | 1 | 2 | nil←1  2→3 |
| 2 | 2 | 3 | nil←1←2  3 |
| 3 | 3 | nil | nil←1←2←3 → head = **3** |

## Approach 3: Recursive
Reverse the rest first, then fix one link:
```
reverse(head):
    if head == nil or head.Next == nil: return head
    newHead = reverse(head.Next)
    head.Next.Next = head   // the node after me should point back to me
    head.Next = nil
    return newHead
```
- **Time:** O(n), **Space:** O(n) recursion stack
