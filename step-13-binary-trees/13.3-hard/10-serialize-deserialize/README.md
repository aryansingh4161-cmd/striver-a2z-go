# Serialize and Deserialize a Binary Tree

**Problem (LeetCode 297):** Turn a tree into a string and back, getting the identical tree.

## Approach: Level order with a null marker
- **Serialize:** BFS; write each node's value, or `#` for a missing child, separated by commas.
  `[1 2 3 null null 4 5]` → `"1,2,3,#,#,4,5,#,#,#,#"`
- **Deserialize:** read the first value as the root; keep a queue of nodes waiting for children; each node takes the next two tokens as its left and right child.

- **Time:** O(n), **Space:** O(n)

Preorder with null markers also works (and is shorter to write recursively): `"1,2,#,#,3,4,#,#,5,#,#"`. Both are in `solution.go`.
