# Preorder, Inorder and Postorder in One Pass

Use a stack of `(node, state)`:
- state 1 → add to **preorder**, set state 2, push the left child (state 1)
- state 2 → add to **inorder**, set state 3, push the right child (state 1)
- state 3 → add to **postorder**, pop

Each node is "visited" three times, exactly mirroring the three points in the recursive function (before left, between, after right).
- **Time:** O(3n) = O(n), **Space:** O(h)
