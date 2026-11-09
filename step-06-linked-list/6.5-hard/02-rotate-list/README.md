# Rotate a Linked List

**Problem (LeetCode 61):** Rotate the list to the **right** by `k` places.

`1 2 3 4 5`, k = 2 → `4 5 1 2 3`

## Approach 1: Move the last node to the front, k times
**O(n × k)**. Too slow when k is big.

## Approach 2: Make a ring, then cut (optimal)
1. Find the length `n` and the tail.
2. `k = k % n` (rotating by n gives the same list). If 0 → return.
3. Link `tail.Next = head` (now it's a circle).
4. The new tail is the node at position `n − k` (1-based). Walk there, set `newHead = newTail.Next`, cut `newTail.Next = nil`.

- **Time:** O(n), **Space:** O(1)
