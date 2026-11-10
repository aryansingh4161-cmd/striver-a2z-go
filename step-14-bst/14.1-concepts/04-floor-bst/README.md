# Floor in a BST

**Problem:** Largest value `<= x`, or −1.

BST `[8 4 12 2 6 10 14]`, x = 11 → 10, x = 1 → −1

## Approach
Mirror of [ceil](../03-ceil-bst/):
- `node == x` → floor is `x`.
- `node < x` → candidate; record it and go **right** for a bigger one.
- `node > x` → go **left**.

- **Time:** O(h), **Space:** O(1)
