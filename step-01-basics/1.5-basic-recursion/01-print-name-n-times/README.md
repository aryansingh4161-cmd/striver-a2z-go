# Print Name N Times

**Problem:** Print a name `n` times using recursion (no loops).

## Approach
Call `f(name, n)`. Print once, then call `f(name, n-1)`. Stop when `n == 0`.

```
f(name, n):
    if n == 0: return
    print(name)
    f(name, n - 1)
```
- **Time:** O(n) calls
- **Space:** O(n) stack depth

## Recursion trace (n = 3)
`f(3)` prints → `f(2)` prints → `f(1)` prints → `f(0)` returns.
