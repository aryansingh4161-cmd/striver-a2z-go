# Detect a Loop (Cycle) in a Linked List

**Problem (LeetCode 141):** Does the list have a cycle (some node's `Next` points back to an earlier node)?

## Approach 1: Hash set of visited nodes
Store each **node pointer** (not value; values can repeat) in a set. Seeing a node twice → cycle.
- **Time:** O(n), **Space:** O(n)

## Approach 2: Floyd's tortoise and hare (optimal)
`slow` moves 1, `fast` moves 2.
- No cycle → `fast` hits `nil`.
- Cycle → once both are inside the loop, `fast` gains **one step per move** on `slow`, so the gap shrinks by 1 every time and they must meet.

```
slow = fast = head
while fast != nil and fast.Next != nil:
    slow = slow.Next; fast = fast.Next.Next
    if slow == fast: return true
return false
```
- **Time:** O(n), **Space:** O(1)

Compare **pointers** (`slow == fast`), not values.
