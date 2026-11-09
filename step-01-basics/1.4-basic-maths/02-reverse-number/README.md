# Reverse a Number

**Problem:** Reverse the digits of `n`. (LeetCode 7 also asks you to return 0 if the result overflows a 32-bit integer.)

| Input | Output |
|---|---|
| 123 | 321 |
| 1200 | 21 |
| -45 | -54 |

## Intuition
Take digits from the **end** of `n` and add them to the **end** of `rev`.
Multiplying `rev` by 10 shifts its digits left to make room.

## Approach
```
rev = 0
while n > 0:
    digit = n % 10
    rev = rev * 10 + digit
    n = n / 10
```
- **Time:** O(log₁₀ n)
- **Space:** O(1)

## Dry run (n = 123)
| n | digit | rev |
|---|---|---|
| 123 | 3 | 3 |
| 12 | 2 | 32 |
| 1 | 1 | **321** |

## Edge cases
- Trailing zeros vanish: 1200 → 21 (that's correct).
- Negative: reverse `-n`, then put the sign back.
- **32-bit overflow (LeetCode 7):** reverse of 1534236469 is bigger than 2³¹−1, so return 0. See `Reverse32`.
