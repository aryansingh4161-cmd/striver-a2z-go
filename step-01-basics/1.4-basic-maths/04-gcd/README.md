# GCD / HCF of Two Numbers

**Problem:** Find the greatest number that divides both `a` and `b`.

| Input | Output |
|---|---|
| 12, 18 | 6 |
| 7, 13 | 1 |
| 20, 0 | 20 |

## Approach 1: Brute force
Try every `i` from `min(a,b)` down to 1. The first `i` that divides both is the answer.
- **Time:** O(min(a, b)), too slow for large numbers.

## Approach 2: Euclidean algorithm (optimal)
**Key fact:** `gcd(a, b) = gcd(b, a % b)`, and `gcd(a, 0) = a`.

Why it works: any number that divides `a` and `b` also divides `a − b`, and so also `a % b` (which is `a` minus `b` several times). So the common divisors don't change, but the numbers get smaller fast.

```
while b != 0:
    a, b = b, a % b
return a
```
- **Time:** O(log(min(a, b)))
- **Space:** O(1)

## Dry run (12, 18)
| a | b | a % b |
|---|---|---|
| 12 | 18 | 12 |
| 18 | 12 | 6 |
| 12 | 6 | 0 |
| 6 | 0 | stop → **6** |

Notice: if `a < b`, the first step simply swaps them.

## Bonus: LCM
`lcm(a, b) = a / gcd(a, b) * b` (divide first to avoid overflow).
