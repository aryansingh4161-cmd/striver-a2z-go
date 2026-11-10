# Divide Two Integers (without *, /, %)

**Problem (LeetCode 29):** Integer division (truncate toward zero). The result must fit in 32 bits: the only overflow is `−2³¹ / −1`, which returns `2³¹ − 1`.

`10 / 3` → 3, `7 / −3` → −2

## Approach 1: Repeated subtraction
Subtract the divisor until you can't. **O(quotient)**: for `2³¹ / 1` that's 2 billion steps.

## Approach 2: Subtract the biggest doubled divisor (optimal)
Every quotient is a sum of powers of 2. Work with absolute values:
```
while dividend >= divisor:
    find the biggest k with divisor << k <= dividend
    dividend -= divisor << k
    quotient += 1 << k
```
Then apply the sign (negative if exactly one input was negative).

- **Time:** O((log n)²), or O(log n) going from the biggest shift down
- **Space:** O(1)

## Dry run 22 / 3
- biggest: 3·4 = 12 ≤ 22 (3·8 = 24 too big) → q += 4, left 10
- biggest: 3·2 = 6 ≤ 10 → q += 2, left 4
- biggest: 3·1 = 3 ≤ 4 → q += 1, left 1
- 1 < 3 → stop. **q = 7** ✓
