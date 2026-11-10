# Generate Parentheses

**Problem (LeetCode 22):** Return all valid combinations of `n` pairs of parentheses.

n = 3 → `((())) (()()) (())() ()(()) ()()()`

## Approach 1: Generate all 2²ⁿ strings, keep valid ones
**O(2²ⁿ × n)**. Wasteful.

## Approach 2: Backtracking with counts (optimal)
Track `open` and `close` used so far:
- add `(` if `open < n`
- add `)` if `close < open` (never close more than you've opened)
- when length is `2n`, save it

Every string built this way is valid, so nothing is wasted.
- **Time:** O(4ⁿ / √n) (the number of valid strings, the Catalan number, times n), **Space:** O(n) depth

## Recursion tree (n = 2)
```
""
└─ (
   ├─ ((
   │  └─ (()
   │     └─ (())  ✓
   └─ ()
      └─ ()(
         └─ ()()  ✓
```
