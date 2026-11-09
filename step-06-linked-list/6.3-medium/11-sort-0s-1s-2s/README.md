# Sort a Linked List of 0s, 1s and 2s

**Problem:** `1 → 0 → 2 → 1 → 0` → `0 → 0 → 1 → 1 → 2`

## Approach 1: Count and overwrite values
Count 0s, 1s, 2s; walk again writing them back. **Time O(2n), Space O(1)**. Changes **values**, which some interviewers don't allow.

## Approach 2: Three chains (optimal, re-links nodes)
Make three dummy heads: `zero`, `one`, `two`. Walk the list once, attaching each node to the matching chain. Then join: `zero tail → one head → two head`.
Careful: if the `one` chain is empty, the `zero` tail must link straight to the `two` chain.
- **Time:** O(n), one pass, **Space:** O(1)
