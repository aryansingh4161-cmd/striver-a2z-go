# Next Greater Element II (circular array)

**Problem (LeetCode 503):** The array is circular: after the last element comes the first. Find the next greater element for each.

`[1 2 1]` → `[2 −1 2]`

## Approach 1: Brute force
For each `i`, check the next `n − 1` elements using index `(i + j) % n`. **O(n²)**.

## Approach 2: Monotonic stack over a "doubled" array (optimal)
Pretend the array is written twice: `[1 2 1 1 2 1]`. Run the [normal algorithm](../01-next-greater-element/) from index `2n − 1` down to 0, reading `a[i % n]`. Only **record answers** for `i < n`; the second copy just fills the stack with "what comes after the wrap".
- **Time:** O(2n) = O(n), **Space:** O(n)
