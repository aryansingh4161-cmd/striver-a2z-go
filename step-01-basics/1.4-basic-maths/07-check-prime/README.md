# Check Prime

**Problem:** A prime has exactly two divisors: 1 and itself. Check if `n` is prime.

| Input | Output |
|---|---|
| 7 | true |
| 12 | false |
| 1 | false (only one divisor) |
| 2 | true (the only even prime) |

## Approach 1: Count all divisors
Count divisors from 1 to n; prime if count == 2.
- **Time:** O(n)

## Approach 2: Check up to √n (optimal)
If `n` had a divisor bigger than √n, its pair partner would be smaller than √n. So if nothing in `2..√n` divides `n`, nothing else does either.

```
if n < 2: return false
for i = 2; i*i <= n; i++:
    if n % i == 0: return false
return true
```
- **Time:** O(√n)
- **Space:** O(1)

## Dry run (n = 29)
Check i = 2, 3, 4, 5 (6×6 = 36 > 29 → stop). None divide 29 → **prime**.

## Edge cases
- 0 and 1 are **not** prime.
- Negative numbers are not prime.

## Going further
To find **all** primes up to N, use the Sieve of Eratosthenes. It comes later in the sheet (maths section of bit manipulation).
