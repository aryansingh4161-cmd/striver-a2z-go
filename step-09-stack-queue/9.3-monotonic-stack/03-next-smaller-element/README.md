# Next Smaller Element (and Previous Smaller)

**Problem:** For each element, the first element to the **right** that is **smaller**. −1 if none.

`[4 8 5 2 25]` → `[2 5 2 −1 −1]`

## Approach: Monotonic stack (optimal)
Mirror of [next greater](../01-next-greater-element/): scan from the right, keep the stack **increasing** from bottom to top, pop everything `>= a[i]`.
- **Time:** O(n), **Space:** O(n)

## The four variants
| Want | Scan from | Pop while top is |
|---|---|---|
| next greater | right | `<= a[i]` |
| next smaller | right | `>= a[i]` |
| previous greater | left | `<= a[i]` |
| previous smaller | left | `>= a[i]` |

The "previous smaller / next smaller" pair, usually storing **indexes** instead of values, powers histogram, sum of subarray minimums and more. Both are in `solution.go`.
