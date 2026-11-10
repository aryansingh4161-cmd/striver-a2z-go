# Count Number of Nice Subarrays

**Problem (LeetCode 1248):** A subarray is nice if it has exactly `k` **odd** numbers. Count them.

`[1 1 2 1 1]`, k = 3 → 2

## Key insight
Replace every number with `1` if odd, `0` if even. Now it's exactly [binary subarrays with sum](../05-binary-subarrays-with-sum/) with goal `k`.

`exactly(k) = atMost(k) − atMost(k − 1)`, using `a[i] % 2` (or `a[i] & 1`) as the value.
- **Time:** O(n), **Space:** O(1)

Recognising "this is the same problem in disguise" is the whole trick.
