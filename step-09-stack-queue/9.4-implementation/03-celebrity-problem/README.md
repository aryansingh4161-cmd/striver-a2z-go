# The Celebrity Problem

**Problem:** `n` people at a party. `M[i][j] = 1` means person `i` knows person `j`. A **celebrity** is known by everyone else and knows nobody. Return the celebrity's index, or −1.

```
M = 0 1 0
    0 0 0     → 1 (everyone knows 1, and 1 knows nobody)
    0 1 0
```

## Approach 1: Count knows / known-by
For every person, count how many they know and how many know them. Celebrity: knows 0, known by n−1. **Time O(n²), Space O(n)**.

## Approach 2: Elimination with two pointers (optimal)
Compare persons `a` (from the top) and `b` (from the bottom):
- if `a` knows `b` → `a` can't be the celebrity (celebrities know nobody) → `a++`
- else → `b` can't be the celebrity (not known by `a`) → `b--`

Each question eliminates one person; after `n − 1` questions one candidate is left. **Verify** it (row all 0, column all 1 except itself), because there may be no celebrity at all.
- **Time:** O(n), **Space:** O(1)

(The classic stack version pushes everyone, pops two at a time, and pushes back the survivor: same idea, O(n) space.)
