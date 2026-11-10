# All Divisors of a Number

Same problem as [Step 1.4](../../../step-01-basics/1.4-basic-maths/06-all-divisors/). Recap:

Divisors come in pairs `(i, n/i)`. Loop `i` while `i*i <= n`; add both `i` and `n/i` (once if they're equal). Sort at the end.
- **Time:** O(√n + d log d), **Space:** O(d)

36 → `1 2 3 4 6 9 12 18 36`
