# Fruit Into Baskets

**Problem (LeetCode 904):** `fruits[i]` is the fruit type of tree `i`. You have **2 baskets**, each holds one type (unlimited amount). Start at any tree and pick one fruit from every tree going right; stop when a fruit doesn't fit. Maximum fruits?

`[1 2 3 2 2]` → 4 (`2 3 2 2`)

## Rephrase
"Longest subarray with **at most 2 distinct values**."

## Approach: Sliding window with a count map (optimal)
Expand `r`, `count[f[r]]++`. While the map has more than 2 keys, decrease `count[f[l]]` (delete the key at 0), `l++`.
- **Time:** O(n), **Space:** O(1) (at most 3 keys)

This is [longest substring with at most K distinct characters](../../10.2-hard/01-longest-substring-k-distinct/) with K = 2.
