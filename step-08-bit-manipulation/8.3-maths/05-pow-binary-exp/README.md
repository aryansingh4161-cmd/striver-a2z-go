# Pow(x, n) with Binary Exponentiation

Covered in detail in [Step 7.1](../../../step-07-recursion/7.1-strong-base/02-pow/). This version focuses on the **bit view**, plus the **modular** form used in many problems.

## The bit view
Write `n` in binary, e.g. `13 = 1101` = 8 + 4 + 1. So `x¹³ = x⁸ · x⁴ · x¹`.
Walk through the bits of n from the right while squaring x each step (x, x², x⁴, x⁸, …). Multiply the result by the current power whenever the bit is 1.
```
result = 1
while n > 0:
    if n & 1: result *= x
    x *= x
    n >>= 1
```
- **Time:** O(log n), **Space:** O(1)

## Modular version
Problems often ask for `xⁿ mod 10⁹+7`. Take `% mod` after every multiplication so numbers never overflow (both factors are < 10⁹+7, so their product fits in Go's 64-bit `int`).
