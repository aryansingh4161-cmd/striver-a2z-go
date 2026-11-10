# Iterative Postorder with Two Stacks

**Trick:** Postorder is Left, Right, Root. **Reversed**, that's Root, Right, Left, which is preorder with the children swapped (easy with one stack).
1. Stack 1: pop node, push it onto stack 2, push **Left then Right** onto stack 1.
2. Stack 2 now holds nodes in Root-Right-Left order; popping it gives Left-Right-Root.

- **Time:** O(n), **Space:** O(n)
