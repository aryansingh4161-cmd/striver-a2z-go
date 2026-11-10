# Check if an Array Represents a Min-Heap

**Problem:** Is the array a valid min-heap (every parent ≤ its children)?

`[10 20 30 21 23]` → true, `[10 20 30 25 15]` → false (20 > 15)

## Approach
Only **non-leaf** nodes have children: indexes `0 .. n/2 − 1`. For each, check `a[i] <= a[2i+1]` and (if it exists) `a[i] <= a[2i+2]`.
- **Time:** O(n), **Space:** O(1)

The "complete tree" shape is automatic: a slice has no gaps.
