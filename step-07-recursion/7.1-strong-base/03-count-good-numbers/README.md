# Count Good Numbers

**Problem (LeetCode 1922):** A digit string of length `n` is **good** if digits at **even** indexes are even (0,2,4,6,8: 5 choices) and digits at **odd** indexes are prime (2,3,5,7: 4 choices). Count good strings of length `n`, modulo 10⁹+7. `n` can be up to 10¹⁵.

n = 1 → 5, n = 4 → 400

## Counting
- Even positions: `ceil(n/2)` = `(n+1)/2` → `5^((n+1)/2)` ways
- Odd positions: `n/2` → `4^(n/2)` ways
- Answer = product, mod 10⁹+7.

With n = 10¹⁵ you can't multiply one by one. Use **fast modular power** (same idea as [Pow(x, n)](../02-pow/), taking `% MOD` after every multiplication).
- **Time:** O(log n), **Space:** O(1)
