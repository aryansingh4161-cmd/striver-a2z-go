# Children Sum Property

**Problem:** Change node values so that **every node equals the sum of its children** (leaves can be anything). You may only **increase** values (by any amount), never decrease. Return the modified tree.

## Approach: Push down, then pull up (one DFS)
Going **down**: if the children's sum is less than the node, set the children equal to the node's value (that only increases them). Then recurse. Going **back up**: set the node to the sum of its (now final) children.
```
fix(n):
    if n == nil: return
    childSum = left.val + right.val (missing = 0)
    if childSum < n.val: set both existing children to n.val
    fix(left); fix(right)
    if n has children: n.val = left.val + right.val
```
Pushing large values down first guarantees that pulling up never needs to **decrease** anything.
- **Time:** O(n), **Space:** O(h)

There are many valid answers; tests check the property and that no value went down.
