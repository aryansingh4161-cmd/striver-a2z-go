# Reverse a Stack Using Recursion

**Problem:** Reverse a stack using only push/pop and recursion.

Bottom → top: `[1 2 3 4]` → `[4 3 2 1]`

## Approach: two recursive functions
1. **reverse(stack):** pop the top `x`, reverse the rest, then **insert `x` at the bottom**.
2. **insertBottom(stack, x):** if empty, push `x`. Otherwise pop the top `t`, insert `x` at the bottom of the rest, push `t` back.

Same shape as [sort a stack](../04-sort-stack/), only the "where to insert" rule changes.
- **Time:** O(n²), **Space:** O(n) recursion stack
