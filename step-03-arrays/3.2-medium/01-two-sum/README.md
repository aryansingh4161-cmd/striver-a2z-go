# Two Sum

**Problem (LeetCode 1):** Return the indices of two different elements that add up to `target`. Exactly one answer exists.
Variant 1: just return yes/no.

`[2 6 5 8 11]`, target 14 → `[1, 3]` (6 + 8)

## Approach 1: Check all pairs (brute)
Two nested loops. **Time O(n²), Space O(1)**.

## Approach 2: Hash map (optimal for indices)
For each `x`, the partner we need is `target − x`. Check if we've already seen it.
```
seen = map value -> index
for i, x in a:
    if target - x in seen: return [seen[target-x], i]
    seen[x] = i
```
- **Time:** O(n), **Space:** O(n)

Check **before** inserting, so an element is never paired with itself.

## Approach 3: Sort + two pointers (good for yes/no)
Sort, then `l = 0, r = n−1`:
- `sum < target` → `l++` (need bigger)
- `sum > target` → `r--` (need smaller)
- equal → found

**Time O(n log n), Space O(1)** if you may change the array. Sorting loses the original indices, so use it for the yes/no version (or sort index pairs).
