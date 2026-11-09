# Print All Divisors

**Problem:** Return every number that divides `n`, in increasing order.

| Input | Output |
|---|---|
| 36 | 1 2 3 4 6 9 12 18 36 |
| 7 | 1 7 |

## Approach 1: Brute force
Check every `i` from 1 to `n`.
- **Time:** O(n)

## Approach 2: Pairs up to √n (optimal)
Divisors come in pairs: if `i` divides `n`, so does `n / i`.
For 36: (1,36) (2,18) (3,12) (4,9) (6,6).

So loop only while `i * i <= n`, and add both `i` and `n/i`.
Careful: when `i == n/i` (like 6 × 6), add it **once**.

- **Time:** O(√n) to find them, plus O(d log d) to sort the d divisors
- **Space:** O(d) for the answer

## Dry run (n = 36)
| i | 36 % i == 0? | added |
|---|---|---|
| 1 | yes | 1, 36 |
| 2 | yes | 2, 18 |
| 3 | yes | 3, 12 |
| 4 | yes | 4, 9 |
| 5 | no | |
| 6 | yes | 6 (once) |
| 7 | 7×7 > 36 → stop | |

Sort → 1 2 3 4 6 9 12 18 36.

## Tip
Write `i*i <= n` instead of `i <= sqrt(n)`. It avoids calling `sqrt` every loop and has no float rounding problems.
