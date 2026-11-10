# Subarrays with K Different Integers

**Problem (LeetCode 992):** Count subarrays with **exactly** `k` distinct integers.

`[1 2 1 2 3]`, k = 2 → 7

## Approach 1: All subarrays
Extend from each start with a map; count while distinct == k; stop once it exceeds k. **O(n²)**.

## Approach 2: atMost(k) − atMost(k − 1) (optimal)
`atMost(k)`: sliding window with a count map, shrinking while distinct > k, adding `r − l + 1` each step.
Exactly the same trick as [binary subarrays with sum](../../10.1-medium/05-binary-subarrays-with-sum/) and [K distinct characters in Step 5](../../../step-05-strings/5.2-medium/05-count-substrings-k-distinct/).
- **Time:** O(n), **Space:** O(k)

## Why "exactly" doesn't work directly
For a fixed `r`, the valid starts for "exactly k" form a range in the **middle**: too far left gives more than k distinct, too far right gives fewer. One window can't capture both edges. Two "at most" windows can.
