# Leaders in an Array

**Problem:** An element is a **leader** if it is greater than every element to its right. The last element is always a leader. Return leaders in their original order.

`[10 22 12 3 0 6]` → `[22 12 6]`

## Approach 1: Check the right side (brute)
For each element, scan everything to its right. **Time O(n²)**.

## Approach 2: Scan from the right (optimal)
Walk from the right, keeping the **maximum seen so far**. If the current element is bigger than that max, it's a leader.
```
maxRight = -∞
for i = n-1 down to 0:
    if a[i] > maxRight: leader; maxRight = a[i]
reverse the leaders (they were collected right-to-left)
```
- **Time:** O(n), **Space:** O(1) besides the answer

## Dry run `[10 22 12 3 0 6]`
| i | a[i] | maxRight | leader? |
|---|---|---|---|
| 5 | 6 | −∞ | yes → 6 |
| 4 | 0 | 6 | no |
| 3 | 3 | 6 | no |
| 2 | 12 | 6 | yes → 12 |
| 1 | 22 | 12 | yes → 22 |
| 0 | 10 | 22 | no |
