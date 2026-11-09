# Maximum Nesting Depth of Parentheses

**Problem (LeetCode 1614):** Valid parentheses string (with other characters too). Return the maximum depth.

`"(1+(2*3)+((8)/4))+1"` → 3

## Approach (optimal)
Counter `depth`: `+1` on `(`, `−1` on `)`. The answer is the highest value `depth` ever reaches.
- **Time:** O(n), **Space:** O(1)

No stack needed: we only care about **how many** are open, not which ones.
