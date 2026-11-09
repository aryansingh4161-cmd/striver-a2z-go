# Delete the Middle Node

**Problem (LeetCode 2095):** Delete the node at index `n/2` (0-based), i.e. the **second** middle for even lengths.

`1 → 3 → 4 → 7 → 1 → 2 → 6` → delete 7
`1 → 2 → 3 → 4` → delete 3

## Approach 1: Count, walk to n/2 − 1
Two passes. **O(n)**.

## Approach 2: Slow/fast with a head start (optimal)
We need to stop **one node before** the middle. Start `fast` two steps ahead (`fast = head.Next.Next`), then move normally. When `fast` runs out, `slow` is just before the middle.
- **Time:** O(n), **Space:** O(1)

One node → return nil.
