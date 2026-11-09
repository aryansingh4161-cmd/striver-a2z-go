# Print 1 to N

**Problem:** Print numbers 1 to N using recursion.

## Approach 1: Count up
`f(i, n)`: if `i > n` stop; print `i`; call `f(i+1, n)`.

## Approach 2: Backtracking (no extra parameter)
`f(n)`: if `n == 0` stop; **first** call `f(n-1)`, **then** print `n`.

The print happens while the stack unwinds, so the smallest number prints first.

```
f(3) → f(2) → f(1) → f(0) returns
             print 1
      print 2
print 3
```
- **Time:** O(n), **Space:** O(n) stack
