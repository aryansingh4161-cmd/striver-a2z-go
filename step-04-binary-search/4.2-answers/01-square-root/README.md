# Square Root of a Number (floor)

**Problem (LeetCode 69):** Return `floor(√n)` without using a library square root.

| n | answer |
|---|---|
| 36 | 6 |
| 28 | 5 |
| 0 | 0 |

## Approach 1: Linear
Try `i = 1, 2, …` while `i*i <= n`. **O(√n)**.

## Approach 2: Binary search on the answer (optimal)
Answers 0..n look like ✅ ✅ ✅ ❌ ❌ for "`x*x <= n`". We want the **last** ✅:
```
on mid*mid <= n: ans = mid, lo = mid + 1   (try bigger)
else: hi = mid - 1
```
- **Time:** O(log n), **Space:** O(1)

`mid*mid` could overflow for huge n in 32-bit languages; Go's `int` is fine up to n ≈ 9 × 10¹⁸ for this range (mid ≤ ~3 × 10⁹ when hi is capped). The solution caps `hi` at a safe value anyway.
