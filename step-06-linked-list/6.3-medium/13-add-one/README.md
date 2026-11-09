# Add 1 to a Number Represented as a Linked List

**Problem:** Digits are stored most significant first. Add 1.

`1 → 5 → 9` → `1 → 6 → 0`
`9 → 9` → `1 → 0 → 0`

The carry has to move from the **last** digit backwards, but the list only goes forward.

## Approach 1: Reverse, add, reverse back
1. Reverse the list.
2. Add 1 with carry from the (new) head.
3. If a carry remains, append a node `1`.
4. Reverse back.
- **Time:** O(3n), **Space:** O(1)

## Approach 2: Recursion (backtracking carry)
Recurse to the end; the base case returns carry 1 (that's the "+1"). On the way back, each node does `val += carry`, keeps `val % 10`, returns `val / 10`. If the head returns a carry, add a new head `1`.
- **Time:** O(n), **Space:** O(n) recursion stack
