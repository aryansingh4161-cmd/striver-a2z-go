# Prime Factors of a Number

**Problem:** Return the **distinct** prime factors of `n`, in increasing order.

60 → `[2 3 5]` (60 = 2² · 3 · 5)

## Approach 1: Check every i, test if it's prime
For `i` in 2..n: if `i` divides n **and** `i` is prime, add it. **O(n√n)**.

## Approach 2: Divide out factors up to √n (optimal)
```
for i = 2; i*i <= n; i++:
    if n % i == 0:
        add i
        while n % i == 0: n /= i     // remove every copy of i
if n > 1: add n                     // what's left is a prime bigger than √(original)
```
Why don't we need a prime check? By the time we reach `i`, every smaller prime has been divided out of `n`, so if `i` divides `n` it **must** be prime (a composite i would have a smaller prime factor that's already gone).

- **Time:** O(√n), **Space:** O(log n) for the answer

## Dry run 60
| i | n before | divides? | n after | factors |
|---|---|---|---|---|
| 2 | 60 | yes | 15 | 2 |
| 3 | 15 | yes | 5 | 2 3 |
| 4 | 5 | 4·4 > 5 → stop | | |
n = 5 > 1 → add 5 → **[2 3 5]**
