# 1.4 Basic Maths

Two tricks power almost every problem here:

| Trick | Meaning | Example (n = 472) |
|---|---|---|
| `n % 10` | last digit | 2 |
| `n / 10` | remove last digit | 47 |

Loop until `n == 0` to visit every digit. That loop runs **number of digits** times = `O(log₁₀ n)`.

The second idea is the **√n trick**: divisors come in pairs `(i, n/i)`. If `i ≤ √n` then `n/i ≥ √n`. So you only need to check `i` from 1 to √n. Write the condition as `i*i <= n` to avoid floating point.

| # | Problem | Key idea |
|---|---|---|
| 1 | [Count digits](01-count-digits/) | divide by 10 until 0 |
| 2 | [Reverse number](02-reverse-number/) | `rev = rev*10 + digit` |
| 3 | [Palindrome number](03-palindrome-number/) | reverse == original |
| 4 | [GCD](04-gcd/) | Euclid: `gcd(a,b) = gcd(b, a%b)` |
| 5 | [Armstrong number](05-armstrong-number/) | sum of digit^count |
| 6 | [All divisors](06-all-divisors/) | pairs up to √n |
| 7 | [Check prime](07-check-prime/) | no divisor in 2..√n |
