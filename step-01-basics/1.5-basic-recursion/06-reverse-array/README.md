# Reverse an Array (recursion)

**Problem:** Reverse an array in place using recursion.

`[1 2 3 4 5]` → `[5 4 3 2 1]`

## Approach 1: Two pointers
`f(l, r)`: if `l >= r` stop; swap `a[l]` and `a[r]`; call `f(l+1, r-1)`.

## Approach 2: One pointer
The partner of index `i` is `n-1-i`. `f(i)`: if `i >= n/2` stop; swap `a[i]` and `a[n-1-i]`; call `f(i+1)`.

- **Time:** O(n) (n/2 swaps), **Space:** O(n) stack

## Dry run `[1 2 3 4 5]`
| l, r | after swap |
|---|---|
| 0, 4 | 5 2 3 4 1 |
| 1, 3 | 5 4 3 2 1 |
| 2, 2 | stop (middle stays) |

In real code, use a loop or `slices.Reverse(a)`. Recursion uses O(n) extra stack for no gain.
