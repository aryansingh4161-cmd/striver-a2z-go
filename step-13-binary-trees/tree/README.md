# Shared Tree Helpers + Representation

```go
type Node struct {
    Val         int
    Left, Right *Node
}
```

## How trees are written in tests: level order with nulls
LeetCode-style: list nodes level by level, left to right, writing `null` for a missing child.

```
[1, 2, 3, null, 4]

      1
     / \
    2   3
     \
      4
```
In Go we can't put `nil` in an `[]int`, so the helper uses a special value `tree.Null`:
```go
root := tree.FromLevel([]int{1, 2, 3, tree.Null, 4})
```
`tree.ToLevel(root)` does the reverse (trailing nulls removed), handy for checking trees you built or modified.

## Other ways to store a tree
- **Array (heap style):** node `i` has children `2i+1`, `2i+2`. Perfect for complete trees (heaps), wasteful otherwise.
- **Pointers (above):** the normal way for interview problems.

## Types of binary trees
| Type | Rule |
|---|---|
| Full | every node has 0 or 2 children |
| Complete | all levels full except the last, filled left to right |
| Perfect | all leaves at the same depth; 2^h − 1 nodes |
| Balanced | heights of left and right subtrees differ by ≤ 1 at every node |
| Degenerate | every node has one child (it's a linked list) |
