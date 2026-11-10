# Two Numbers That Appear Once (others twice)

**Problem (LeetCode 260):** Exactly two numbers appear once; all others appear twice. Find both.

`[1 2 1 3 2 5]` → `[3 5]`

## Approach 1: Hash map
**Time O(n), Space O(n)**.

## Approach 2: XOR + split by one bit (optimal)
1. XOR everything → `xr = x ^ y` (pairs cancel). Since `x ≠ y`, `xr ≠ 0`.
2. Take **any bit where they differ**: the lowest set bit `bit = xr & −xr`.
3. Split all numbers into two groups by that bit. `x` and `y` land in **different** groups; each pair lands together in one group and cancels.
4. XOR each group → one gives `x`, the other gives `y`.

- **Time:** O(n), **Space:** O(1)

## Dry run `[1 2 1 3 2 5]`
`xr = 3 ^ 5 = 011 ^ 101 = 110`. Lowest set bit = `010`.
Group bit set: 2, 3, 2 → XOR = **3**. Group bit clear: 1, 1, 5 → XOR = **5**.
