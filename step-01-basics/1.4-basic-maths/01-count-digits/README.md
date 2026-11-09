# Count Digits

**Problem:** Given a number `n`, return how many digits it has.

| Input | Output |
|---|---|
| 12345 | 5 |
| 7 | 1 |
| 0 | 1 |

## Intuition
Every `n / 10` removes one digit. Count how many times you can do it before `n` becomes 0.

## Approach 1: Divide by 10 (main approach)
```
count = 0
while n > 0:
    count++
    n = n / 10
```
- **Time:** O(log₁₀ n), one step per digit
- **Space:** O(1)

## Approach 2: Logarithm
Number of digits = `floor(log10(n)) + 1`. One line, but it uses floating point, which can be off by one for very large numbers. Good to know, but prefer the loop.

## Dry run (n = 472)
| n before | count |
|---|---|
| 472 | 1 |
| 47 | 2 |
| 4 | 3 |
| 0 | stop → **3** |

## Edge cases
- `n = 0` → the loop never runs, but 0 has **1** digit. Handle it separately.
- Negative numbers → count digits of `-n`.
