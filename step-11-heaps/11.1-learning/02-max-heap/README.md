# Implement a Max-Heap

Same as the [min-heap](../01-min-heap-from-scratch/), with every comparison flipped: parent ≥ children, the **maximum** is at index 0.

## Two ways in Go
1. **Hand-written**: copy the min-heap and flip `<` to `>`.
2. **`container/heap`**: implement the 5 methods with `Less(i, j) = h[i] > h[j]`.

A third trick from competitive programming: store **negated** values in a min-heap. Works, but is easy to get wrong when you print values. Prefer flipping `Less`.

Both versions 1 and 2 are in `solution.go`.
