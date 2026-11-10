# XOR of Numbers in a Range L..R

**Problem:** Compute `L ^ (L+1) ^ … ^ R`.

## Approach 1: Loop
**O(R − L)**.

## Approach 2: Pattern for 1..n (optimal)
`f(n) = 1 ^ 2 ^ … ^ n` repeats every 4:

| n % 4 | f(n) |
|---|---|
| 0 | n |
| 1 | 1 |
| 2 | n + 1 |
| 3 | 0 |

Check: f(1)=1, f(2)=3, f(3)=0, f(4)=4, f(5)=1, f(6)=7, f(7)=0, f(8)=8 ✓

Then, like prefix sums: `XOR(L..R) = f(R) ^ f(L−1)` (the part `1..L−1` cancels).
- **Time:** O(1)
