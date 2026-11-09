# Middle of the Linked List

**Problem (LeetCode 876):** Return the middle node. With two middles (even length), return the **second** one.

`1 → 2 → 3 → 4 → 5` → 3
`1 → 2 → 3 → 4 → 5 → 6` → 4

## Approach 1: Count, then walk
Count `n`, then walk `n/2` steps. **Time O(n + n/2), Space O(1)**. Two passes.

## Approach 2: Tortoise and hare (optimal, one pass)
`slow` moves 1 step, `fast` moves 2. When `fast` reaches the end, `slow` is halfway.
```
slow = fast = head
while fast != nil and fast.Next != nil:
    slow = slow.Next
    fast = fast.Next.Next
return slow
```
- **Time:** O(n/2), **Space:** O(1)

## Dry run `1 2 3 4 5 6`
| slow | fast |
|---|---|
| 1 | 1 |
| 2 | 3 |
| 3 | 5 |
| 4 | nil (past 6) → stop, middle = **4** |

Want the **first** middle for even lengths (used in palindrome / sort)? Loop while `fast.Next != nil && fast.Next.Next != nil`.
