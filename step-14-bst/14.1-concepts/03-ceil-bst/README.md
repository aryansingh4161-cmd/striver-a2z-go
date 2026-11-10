# Ceil in a BST

**Problem:** Smallest value `>= x`, or −1.

BST `[8 4 12 2 6 10 14]`, x = 11 → 12, x = 6 → 6, x = 15 → −1

## Approach: Walk down, remembering the best candidate
- `node == x` → that's the ceil.
- `node > x` → this node **could** be the ceil; record it and go **left** for a smaller one.
- `node < x` → too small; go **right**.

- **Time:** O(h), **Space:** O(1)

Same "save and keep searching" idea as [ceil in a sorted array](../../../step-04-binary-search/4.1-1d/05-floor-and-ceil/).
