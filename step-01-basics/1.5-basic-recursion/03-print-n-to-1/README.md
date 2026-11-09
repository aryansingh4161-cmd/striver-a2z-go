# Print N to 1

**Problem:** Print numbers N down to 1 using recursion.

## Approach 1: Count down
`f(n)`: if `n == 0` stop; print `n`; call `f(n-1)`.

## Approach 2: Backtracking
`f(i, n)`: if `i > n` stop; call `f(i+1, n)` **first**, then print `i`. The largest number prints first because it is deepest in the stack.

- **Time:** O(n), **Space:** O(n) stack

Compare with [Print 1 to N](../02-print-1-to-n/): same code, the print just moves before or after the call.
