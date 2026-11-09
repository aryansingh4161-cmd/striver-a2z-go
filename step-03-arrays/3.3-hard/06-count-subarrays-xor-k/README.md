# Count Subarrays with XOR K

**Problem:** Count subarrays whose XOR equals `k`.

`[4 2 2 6 4]`, k = 6 → 4

## XOR refresher
- `x ^ x = 0`, `x ^ 0 = x`
- If `A ^ B = C` then `A = C ^ B`.

## Brute
All subarrays with a running XOR. **O(n²)**.

## Optimal: prefix XOR + frequency map
Let `xr` = XOR of `a[0..i]`. A subarray ending at `i` with XOR `k` exists when an earlier prefix XOR `x` satisfies `x ^ k = xr`, i.e. **`x = xr ^ k`**.
Count how many times `xr ^ k` appeared before. Start with `{0: 1}` (empty prefix).

Exactly the same shape as [count subarrays with sum K](../../3.2-medium/13-count-subarrays-sum-k/), with `^` instead of `−`.

- **Time:** O(n), **Space:** O(n)

## Dry run `[4 2 2 6 4]`, k = 6
| x | xr | xr^k | seen[xr^k] | ans |
|---|---|---|---|---|
| 4 | 4 | 2 | 0 | 0 |
| 2 | 6 | 0 | 1 | 1 |
| 2 | 4 | 2 | 0 | 1 |
| 6 | 2 | 4 | 2 | 3 |
| 4 | 6 | 0 | 1 | **4** |
