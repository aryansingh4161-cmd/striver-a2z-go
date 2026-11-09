# Odd Even Linked List

**Problem (LeetCode 328):** Group the nodes at **odd positions** (1st, 3rd, 5th…) first, then the even positions, keeping relative order. O(1) extra space.

`1 → 2 → 3 → 4 → 5` → `1 → 3 → 5 → 2 → 4`

(Positions, not values!)

## Approach 1: Copy values (brute)
Collect odd-position values, then even-position values, write them back. **Space O(n)**.

## Approach 2: Re-link in place (optimal)
Two chains: `odd` starts at node 1, `even` at node 2. Remember `evenHead`.
```
while even != nil and even.Next != nil:
    odd.Next = even.Next;  odd = odd.Next
    even.Next = odd.Next;  even = even.Next
odd.Next = evenHead      // attach evens after odds
```
- **Time:** O(n), **Space:** O(1)
