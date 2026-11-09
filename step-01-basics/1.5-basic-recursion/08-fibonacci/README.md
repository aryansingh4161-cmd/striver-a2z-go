# Fibonacci Number

**Problem:** `F(0) = 0`, `F(1) = 1`, `F(n) = F(n-1) + F(n-2)`. Return `F(n)`.

Sequence: 0 1 1 2 3 5 8 13 21 …

## Approach 1: Plain recursion (multiple calls)
```
fib(n):
    if n <= 1: return n
    return fib(n-1) + fib(n-2)
```
Each call makes **two** calls → about **O(2ⁿ)** time. `fib(45)` takes seconds.
- **Space:** O(n) (the deepest path)

See the recursion tree in the [recursion notes](../README.md): `fib(2)` is computed again and again. That repeated work is exactly what **Dynamic Programming** (Step 16) removes.

## Approach 2: Loop with two variables (optimal)
Keep only the last two values.
```
prev2, prev1 = 0, 1
repeat n-1 times:
    cur = prev1 + prev2
    prev2, prev1 = prev1, cur
```
- **Time:** O(n), **Space:** O(1)

## Dry run (n = 5)
| step | prev2 | prev1 |
|---|---|---|
| start | 0 | 1 |
| 1 | 1 | 1 |
| 2 | 1 | 2 |
| 3 | 2 | 3 |
| 4 | 3 | **5** |
