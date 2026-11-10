# Step 13: Binary Trees

> Status: ✅ traversals and medium done · ⏳ hard problems coming in the next batch

Shared helpers live in [`tree/`](tree/): `tree.FromLevel([]int{1, 2, 3, tree.Null, 4})` builds a tree from LeetCode-style level order.

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Introduction and representation](tree/) | level-order input with nulls |
| 2 | [Preorder (recursive)](13.1-traversals/01-recursive-traversals/) | root, left, right |
| 3 | [Inorder (recursive)](13.1-traversals/01-recursive-traversals/) | left, root, right |
| 4 | [Postorder (recursive)](13.1-traversals/01-recursive-traversals/) | left, right, root |
| 5 | [Level order](13.1-traversals/02-level-order/) | queue + level size |
| 6 | [Iterative preorder](13.1-traversals/03-iterative-preorder/) | push right, then left |
| 7 | [Iterative inorder](13.1-traversals/04-iterative-inorder/) | dive left, visit, go right |
| 8 | [Postorder with 2 stacks](13.1-traversals/05-postorder-two-stacks/) | reverse of root-right-left |
| 9 | [Postorder with 1 stack](13.1-traversals/06-postorder-one-stack/) | remember last visited |
| 10 | [Pre, in, post in one pass](13.1-traversals/07-all-traversals-one-pass/) | (node, state) stack |

### 13.2 Medium
| # | Problem | Key idea |
|---|---|---|
| 1 | [Maximum depth](13.2-medium/01-max-depth/) | 1 + max(left, right) |
| 2 | [Check balanced](13.2-medium/02-check-balanced/) | height returns −1 when unbalanced |
| 3 | [Diameter](13.2-medium/03-diameter/) | best = max(left + right) |
| 4 | [Maximum path sum](13.2-medium/04-max-path-sum/) | drop negative branches |
| 5 | [Identical trees](13.2-medium/05-identical-trees/) | compare recursively |
| 6 | [Zigzag level order](13.2-medium/06-zigzag-traversal/) | fill level from the back |
| 7 | [Boundary traversal](13.2-medium/07-boundary-traversal/) | left edge + leaves + reversed right edge |
| 8 | [Vertical order](13.2-medium/08-vertical-order/) | (col, row, val) then sort |
| 9 | [Top view](13.2-medium/09-top-view/) | BFS, first per column |
| 10 | [Bottom view](13.2-medium/10-bottom-view/) | BFS, last per column |
| 11 | [Left / right view](13.2-medium/11-left-right-view/) | DFS, first node at each new depth |
| 12 | [Symmetric tree](13.2-medium/12-symmetric-tree/) | mirror with crossed children |

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
