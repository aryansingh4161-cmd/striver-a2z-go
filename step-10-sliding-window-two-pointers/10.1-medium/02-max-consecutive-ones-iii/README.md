# Max Consecutive Ones III

**Problem (LeetCode 1004):** Binary array; you may flip at most `k` zeros to ones. Return the longest run of ones possible.

`[1 1 1 0 0 0 1 1 1 1 0]`, k = 2 → 6

## Rephrase
"Longest subarray containing **at most k zeros**." That's a window condition.

## Approach 1: All subarrays
Count zeros while extending; stop when zeros > k. **O(n²)**.

## Approach 2: Sliding window (optimal)
Expand `r`; count zeros. While `zeros > k`, move `l` right (subtract if `a[l]` was 0). Update the best length.
- **Time:** O(2n), **Space:** O(1)

## Approach 3: Window that never shrinks (slightly faster)
Use `if` instead of `while`: when the window becomes invalid, slide it by one instead of shrinking. Its size never decreases, so at the end `r − l + 1` is the answer. Still O(n), just fewer steps. Both are in the code.
