# Prime Factorisation of Many Numbers (Smallest Prime Factor Sieve)

**Problem:** Answer many queries "give the full prime factorisation of x" (x ≤ 10⁶), fast.

12 → `[2 2 3]`, 97 → `[97]`

## Approach 1: Trial division per query
O(√x) each → O(Q√x) total.

## Approach 2: SPF sieve (optimal for many queries)
Precompute `spf[x]` = the **smallest prime factor** of every x up to N (a sieve variant):
```
spf[i] = i for all i
for p = 2; p*p <= N; p++:
    if spf[p] == p:                 // p is prime
        for m = p*p; m <= N; m += p:
            if spf[m] == m: spf[m] = p
```
Then factorise `x` by repeatedly dividing by its smallest prime factor:
```
while x > 1: add spf[x]; x /= spf[x]
```
Each step at least halves x → **O(log x)** per query.

- **Precompute:** O(N log log N), **Query:** O(log x), **Space:** O(N)
