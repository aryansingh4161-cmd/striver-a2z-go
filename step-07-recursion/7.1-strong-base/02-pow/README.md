# Pow(x, n)

**Problem (LeetCode 50):** Compute `xⁿ` where `n` can be negative.

`2.0, 10` → 1024, `2.0, −2` → 0.25

## Approach 1: Multiply n times
**O(n)**. Too slow for n ≈ 2³¹.

## Approach 2: Fast power / binary exponentiation (optimal)
Use: `xⁿ = (x²)^(n/2)` when n is even, and `xⁿ = x · xⁿ⁻¹` when n is odd.
Each step halves `n` → **O(log n)**.

Recursive form:
```
pow(x, n):
    if n == 0: return 1
    half = pow(x, n / 2)
    if n is even: return half * half
    else: return half * half * x
```
For negative `n`: compute `pow(x, −n)` and return `1 / result`.

Iterative form (O(1) space): while `n > 0`: if `n` is odd, `result *= x`; then `x *= x`, `n /= 2`.

- **Time:** O(log n), **Space:** O(log n) recursive / O(1) iterative

## Dry run 2¹⁰
`2¹⁰ = (2⁵)²`, `2⁵ = 2·(2²)²`, `2² = (2¹)²`, `2¹ = 2·(2⁰)²` → only 4 levels instead of 10 multiplications.
