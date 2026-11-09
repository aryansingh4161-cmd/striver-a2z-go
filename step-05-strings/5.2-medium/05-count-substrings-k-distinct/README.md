# Count Substrings with Exactly K Distinct Characters

**Problem:** Count substrings of `s` (lowercase) that contain exactly `k` distinct characters.

`"pqpqs"`, k = 2 → 7 (`pq, pqp, pqpq, qp, qpq, pq, qs`)

## Approach 1: All substrings (brute)
Fix start `i`, extend `j`, keep a `[26]` count and a distinct counter. Count when distinct == k; stop extending once distinct > k.
- **Time:** O(n²), **Space:** O(1)

## Approach 2: atMost(k) − atMost(k−1) (optimal)
"Exactly k" is awkward for a sliding window, but "**at most** k" is easy:
- expand `r`, shrink `l` while distinct > k,
- every valid window ending at `r` contributes `r − l + 1` substrings (all those starting from `l..r`).

Then `exactly(k) = atMost(k) − atMost(k−1)`.
- **Time:** O(n), **Space:** O(1)

This trick is covered in depth in [Step 10](../../../step-10-sliding-window-two-pointers/).
