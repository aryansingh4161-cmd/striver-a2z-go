# Step 13: Binary Trees

> Status: theory ready, problems coming ⏳

## Words to know
- **Node, root, leaf** (no children), **parent, child, subtree**.
- **Height/depth:** number of levels on the longest root-to-leaf path.
- **Full** binary tree: every node has 0 or 2 children.
- **Complete:** all levels full except maybe the last, which is filled left to right (heaps are complete).
- **Perfect:** all leaves on the same level.
- **Balanced:** for every node, left and right heights differ by at most 1.

```go
type TreeNode struct {
    Val         int
    Left, Right *TreeNode
}
```

## Traversals (learn these until they're automatic)
| Name | Order | Use |
|---|---|---|
| Preorder | Root, Left, Right | copy a tree, serialise |
| Inorder | Left, Root, Right | gives **sorted** order in a BST |
| Postorder | Left, Right, Root | delete a tree, compute from children up (height, diameter) |
| Level order (BFS) | level by level with a queue | views, width, zigzag |

```go
func inorder(root *TreeNode, out *[]int) {
    if root == nil { return }
    inorder(root.Left, out)
    *out = append(*out, root.Val)
    inorder(root.Right, out)
}
```

Level order:
```go
q := []*TreeNode{root}
for len(q) > 0 {
    size := len(q)              // nodes on this level
    for i := 0; i < size; i++ {
        node := q[0]; q = q[1:]
        if node.Left != nil  { q = append(q, node.Left) }
        if node.Right != nil { q = append(q, node.Right) }
    }
}
```

## The one idea behind most tree problems
**Ask each subtree for an answer, combine at the node.** Height = `1 + max(height(left), height(right))`.
Many "hard" problems (diameter, max path sum, balanced check) are the height function plus one extra line that updates a global answer.

## Other key ideas
- **Vertical / top / bottom view:** BFS while tracking each node's column (`left = col−1`, `right = col+1`).
- **Root-to-node path / LCA:** recurse; a node is the LCA when the targets are found on different sides.
- **Building a tree:** inorder + preorder (or postorder) uniquely define a tree. Use a map from value → inorder index.
- **Morris traversal:** inorder in O(1) space by temporarily linking each node's in-order predecessor back to it.

## Problems

### 13.1 Traversals
| # | Problem | Status |
|---|---|---|
| 1 | Introduction and representation | ⏳ |
| 2 | Preorder (recursive) | ⏳ |
| 3 | Inorder (recursive) | ⏳ |
| 4 | Postorder (recursive) | ⏳ |
| 5 | Level order | ⏳ |
| 6 | Iterative preorder | ⏳ |
| 7 | Iterative inorder | ⏳ |
| 8 | Postorder with 2 stacks | ⏳ |
| 9 | Postorder with 1 stack | ⏳ |
| 10 | Pre, in, post in one traversal | ⏳ |

### 13.2 Medium
| # | Problem | Status |
|---|---|---|
| 1 | Maximum depth (height) | ⏳ |
| 2 | Check if balanced | ⏳ |
| 3 | Diameter | ⏳ |
| 4 | Maximum path sum | ⏳ |
| 5 | Identical trees | ⏳ |
| 6 | Zigzag level order | ⏳ |
| 7 | Boundary traversal | ⏳ |
| 8 | Vertical order traversal | ⏳ |
| 9 | Top view | ⏳ |
| 10 | Bottom view | ⏳ |
| 11 | Left / right view | ⏳ |
| 12 | Symmetric tree | ⏳ |

### 13.3 Hard
| # | Problem | Status |
|---|---|---|
| 1 | Root to node path | ⏳ |
| 2 | Lowest common ancestor | ⏳ |
| 3 | Maximum width | ⏳ |
| 4 | Children sum property | ⏳ |
| 5 | All nodes at distance K | ⏳ |
| 6 | Minimum time to burn the tree | ⏳ |
| 7 | Count nodes in a complete tree | ⏳ |
| 8 | Build tree from preorder + inorder | ⏳ |
| 9 | Build tree from postorder + inorder | ⏳ |
| 10 | Serialise and deserialise | ⏳ |
| 11 | Morris inorder traversal | ⏳ |
| 12 | Morris preorder traversal | ⏳ |
| 13 | Flatten binary tree to linked list | ⏳ |
