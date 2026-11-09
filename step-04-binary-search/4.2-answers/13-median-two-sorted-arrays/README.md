# Median of Two Sorted Arrays

**Problem (LeetCode 4):** Two sorted arrays of sizes n1 and n2. Return the median of all elements combined, in O(log(min(n1, n2))).

`[1 3]`, `[2]` → 2.0
`[1 2]`, `[3 4]` → 2.5

## Approach 1: Merge (brute)
Merge like merge sort, take the middle. **Time O(n1 + n2), Space O(n1 + n2)**.

## Approach 2: Merge without storing (better)
Same walk with two pointers, but only remember the values at the middle positions. **Time O(n1 + n2), Space O(1)**.

## Approach 3: Binary search on the partition (optimal)
Cut both arrays so the **left side** holds exactly `half = (n1 + n2 + 1) / 2` elements:
- take `cut1` elements from array 1 and `cut2 = half − cut1` from array 2.

Let `l1, r1` be the elements just left/right of cut 1, and `l2, r2` for cut 2 (use −∞ / +∞ when a side is empty).

The cut is **correct** when everything on the left ≤ everything on the right:
`l1 <= r2` **and** `l2 <= r1`.

- If `l1 > r2` → took too many from array 1 → move `cut1` left (`hi = cut1 − 1`).
- If `l2 > r1` → took too few → move right (`lo = cut1 + 1`).

When correct:
- total odd → median = `max(l1, l2)`
- total even → median = `(max(l1, l2) + min(r1, r2)) / 2`

Binary search over `cut1 ∈ [0, n1]` on the **smaller** array.
- **Time:** O(log(min(n1, n2))), **Space:** O(1)

## Dry run `[1 3 4 7 10 12]` and `[2 3 6 15]`
Total 10, half = 5. Search on the smaller array A = `[2 3 6 15]`; B = `[1 3 4 7 10 12]`. `cutA` ranges over 0..4.

| cutA | cutB = 5 − cutA | lA, rA | lB, rB | lA ≤ rB? | lB ≤ rA? | action |
|---|---|---|---|---|---|---|
| 2 | 3 | 3, 6 | 4, 7 | ✓ | ✓ | correct cut |

Left side: `2 3 | 1 3 4`, right side: `6 15 | 7 10 12`.
Median = (max(3, 4) + min(6, 7)) / 2 = (4 + 6) / 2 = **5.0**
