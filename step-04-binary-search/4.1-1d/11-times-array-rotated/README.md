# How Many Times Has the Array Been Rotated?

**Problem:** A sorted array (distinct values) was rotated **right** k times. Find k.

`[4 5 6 7 0 1 2]` → 4

## Key insight
The number of rotations = the **index of the minimum element**. (After k right-rotations, the original first element, the minimum, sits at index k.)

So run [minimum in rotated sorted array](../10-min-in-rotated-sorted-array/), but track the **index** of the minimum instead of its value.
- **Time:** O(log n), **Space:** O(1)
