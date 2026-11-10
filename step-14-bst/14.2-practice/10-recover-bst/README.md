# Recover a BST (two nodes swapped)

**Problem (LeetCode 99):** Exactly two nodes of a BST had their values swapped by mistake. Fix it without changing the structure.

## Approach 1: Inorder, sort, write back
**O(n log n)**, O(n) space.

## Approach 2: Find the two "drops" in inorder (optimal)
Inorder of a valid BST is increasing. Swapping two values creates either:
- **two drops** (non-adjacent swap): `1 [6] 3 4 5 [2] 7` → first drop at 6>3, second at 5>2. Swap the **first element of the first drop** (6) with the **second element of the second drop** (2).
- **one drop** (adjacent swap): `1 2 [4 3] 5` → swap the two elements of that drop.

Track `prev` during inorder; on `prev.val > cur.val`: if it's the first drop, set `first = prev`, `middle = cur`; else set `last = cur`. At the end swap `first` with `last` if found, else with `middle`.
- **Time:** O(n), **Space:** O(h) (O(1) with Morris inorder)
