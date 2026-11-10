# Power Set Using Bit Manipulation

**Problem:** Return all subsets of `n` distinct elements.

## Approach
Each number `mask` from `0` to `2ⁿ − 1` describes one subset: **bit `i` set → include `a[i]`**.

| mask | binary (a=[x y z]) | subset |
|---|---|---|
| 0 | 000 | [] |
| 1 | 001 | [x] |
| 2 | 010 | [y] |
| 3 | 011 | [x y] |
| … | … | … |
| 7 | 111 | [x y z] |

- **Time:** O(2ⁿ × n), **Space:** O(1) besides the answer

Compare with the recursive version in [Step 7](../../../step-07-recursion/7.2-subsequences/03-power-set/).
