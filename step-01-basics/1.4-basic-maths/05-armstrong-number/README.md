# Armstrong Number

**Problem:** An Armstrong number equals the sum of its digits, each raised to the power of the number of digits. Check if `n` is one.

| Input | Why | Output |
|---|---|---|
| 153 | 1³ + 5³ + 3³ = 1 + 125 + 27 = 153 | true |
| 9474 | 9⁴ + 4⁴ + 7⁴ + 4⁴ = 9474 | true |
| 123 | 1 + 8 + 27 = 36 | false |

## Approach
1. Count the digits → `k` (see [count digits](../01-count-digits/)).
2. Walk the digits with `% 10` and `/ 10`, adding `digit^k` to a sum.
3. Compare the sum with the **original** number.

- **Time:** O(log₁₀ n × k) ≈ O((log n)²), tiny in practice
- **Space:** O(1)

## Note
Many videos use cube (`^3`) for every number. That only works for 3-digit numbers. The correct definition uses the digit count `k`.

## Edge cases
- Every single-digit number (0–9) is an Armstrong number (`d¹ = d`).
