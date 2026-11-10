# Sieve of Eratosthenes

**Problem (LeetCode 204 style):** Find all primes up to `n`.

n = 30 → `2 3 5 7 11 13 17 19 23 29`

## Approach 1: Test each number with the √ check
**O(n√n)**.

## Approach 2: Sieve (optimal)
Start by assuming everything from 2 is prime. For each prime `p`, cross out its multiples: they can't be prime.
```
isPrime[0..n] = true; isPrime[0] = isPrime[1] = false
for p = 2; p*p <= n; p++:
    if isPrime[p]:
        for m = p*p; m <= n; m += p: isPrime[m] = false
```
Two speed-ups:
- Start crossing at `p*p`: smaller multiples like `2p, 3p` were already crossed by smaller primes.
- Stop the outer loop at `√n`: any composite ≤ n has a factor ≤ √n.

- **Time:** O(n log log n) (practically linear), **Space:** O(n)

Use it when you need **many** primality answers, or all primes in a range.
