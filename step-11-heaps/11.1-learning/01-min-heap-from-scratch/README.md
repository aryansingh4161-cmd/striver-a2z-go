# Implement a Min-Heap from Scratch

A **min-heap** is a complete binary tree stored in a slice where every parent ≤ its children. The minimum is always at index 0.

```
index:   0  1  2  3  4  5
value:   1  3  2  7  4  5
tree:          1
             /   \
            3     2
           / \   /
          7   4 5
```
For index `i`: parent `(i−1)/2`, left child `2i+1`, right child `2i+2`.

## Operations
| Operation | How | Time |
|---|---|---|
| `Peek()` | `a[0]` | O(1) |
| `Push(x)` | append at the end, then **sift up**: swap with the parent while smaller | O(log n) |
| `Pop()` | save `a[0]`; move the **last** element to index 0; shrink; **sift down**: swap with the smaller child while bigger | O(log n) |
| `Build(a)` | sift down every non-leaf, from index `n/2 − 1` down to 0 | **O(n)** (not n log n) |

**Why is build O(n)?** Most nodes are near the bottom and sift down only a step or two. Only a few near the top travel far. The total work adds up to about 2n.

**Heapsort:** build a heap, then pop n times → O(n log n), in place.

In real Go code use `container/heap` (template in [docs/go-for-dsa.md](../../../docs/go-for-dsa.md)). This file builds one by hand so you understand what that package does.
