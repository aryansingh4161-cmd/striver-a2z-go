# Sort a Stack Using Recursion

**Problem:** Sort a stack so the **largest** element is on top, using only push/pop/peek and recursion (no loops over the stack, no extra array).

Bottom → top: `[3 1 4 2]` → `[1 2 3 4]`

## Approach: two recursive functions
1. **sort(stack):** pop the top `x`, sort the rest recursively, then **insert `x` in the right place**.
2. **insert(stack, x):** if the stack is empty or `top <= x`, push `x`. Otherwise pop the top `t`, insert `x` into the rest, then push `t` back.

It's [insertion sort](../../../step-02-sorting/03-insertion-sort/) written with the call stack.

- **Time:** O(n²), **Space:** O(n) recursion stack

## Dry run `[3 1 4 2]` (top = 2)
sort pops 2, 4, 1, 3 down to empty, then inserts back:
- insert 3 → `[3]`
- insert 1 → 3 > 1: pop 3, push 1, push 3 → `[1 3]`
- insert 4 → `[1 3 4]`
- insert 2 → pop 4, pop 3, push 2, push 3, push 4 → `[1 2 3 4]`
