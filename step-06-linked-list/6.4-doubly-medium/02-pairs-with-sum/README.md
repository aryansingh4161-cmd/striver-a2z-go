# Find Pairs with a Given Sum in a Sorted DLL

**Problem:** Sorted DLL of distinct values. Return all pairs `(a, b)` with `a + b = target`, `a < b`.

`1 ⇄ 2 ⇄ 4 ⇄ 5 ⇄ 6 ⇄ 8 ⇄ 9`, target 7 → `(1,6) (2,5)`

## Approach 1: Two loops
For each node, scan the rest. **O(n²)**.

## Approach 2: Two pointers from both ends (optimal)
Possible only because a DLL can walk **backwards**. `left` = head, `right` = tail.
- `sum < target` → `left = left.Next`
- `sum > target` → `right = right.Prev`
- equal → record, move both

Stop when they meet or cross (`left == right` or `right.Next == left`).
- **Time:** O(n) (finding the tail + one sweep), **Space:** O(1)
