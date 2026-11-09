# Spiral Traversal of a Matrix

**Problem (LeetCode 54):** Return all elements in spiral order (right → down → left → up, moving inward).

```
1  2  3  4
5  6  7  8     →   1 2 3 4 8 12 11 10 9 5 6 7
9 10 11 12
```

## Approach: Four boundaries (optimal)
Keep `top, bottom, left, right`. Repeat while `top <= bottom && left <= right`:
1. **→** along `top` row from `left` to `right`; then `top++`.
2. **↓** along `right` column from `top` to `bottom`; then `right--`.
3. **←** along `bottom` row from `right` to `left`, **only if `top <= bottom`**; then `bottom--`.
4. **↑** along `left` column from `bottom` to `top`, **only if `left <= right`**; then `left++`.

The two "only if" checks matter for a single remaining row or column; without them you print it twice.

- **Time:** O(n·m), **Space:** O(1) besides the answer
