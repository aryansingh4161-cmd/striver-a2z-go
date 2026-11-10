# Implement a Queue Using Stacks

**Problem (LeetCode 232):** Build FIFO behaviour from stacks.

## Approach 1: Make push expensive
Before pushing, move everything from `s1` to `s2`, push `x` onto `s1`, move everything back. The oldest element ends up on top of `s1`.
- **Push:** O(n), **Pop:** O(1)

## Approach 2: Two stacks, lazy transfer (optimal, amortised O(1))
- `in` stack receives pushes.
- `out` stack serves pops.
- On pop/peek: if `out` is empty, pour **all** of `in` into `out` (this reverses the order, so the oldest is on top). Then pop from `out`.

Each element moves from `in` to `out` **once**, so over many operations each one costs O(1) on average (**amortised O(1)**).

## Picture: push 1, 2, 3, then pop
```
in = [1 2 3], out = []
pop → out empty → pour: in = [], out = [3 2 1] (top is 1) → pop 1 ✓
```
