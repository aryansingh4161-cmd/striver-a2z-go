# Factorial of N

**Problem:** `n! = 1 × 2 × … × n`, and `0! = 1`.

| n | n! |
|---|---|
| 0 | 1 |
| 5 | 120 |

## Approach (functional recursion)
```
fact(n):
    if n <= 1: return 1
    return n * fact(n - 1)
```
- **Time:** O(n), **Space:** O(n) stack

## Watch out
Factorials grow very fast. `20!` ≈ 2.4 × 10¹⁸ is the largest that fits in a 64-bit `int`; `21!` overflows. Problems usually ask for the answer modulo 10⁹+7.
