# Minimum and Maximum in a BST

- **Minimum:** keep going **left** until there's no left child.
- **Maximum:** keep going **right**.

- **Time:** O(h), **Space:** O(1)

These two walks are used inside **delete** (inorder successor = minimum of the right subtree) and the **BST iterator**.
