# Rearrange Array Elements by Sign

**Variant 1 (LeetCode 2149):** Equal number of positives and negatives. Rearrange so signs alternate, starting with positive, keeping the relative order within each sign.

`[3 1 -2 -5 2 -4]` → `[3 -2 1 -5 2 -4]`

**Variant 2:** Counts may differ. Alternate as long as possible, then put the leftovers at the end in their original order.

## Variant 1, brute
Split into `pos` and `neg` slices, then write `pos[i]` at index `2i` and `neg[i]` at `2i+1`.
**Time O(n) + O(n/2), Space O(n)**.

## Variant 1, optimal (one pass)
Positives go to even indexes, negatives to odd. Keep two write pointers:
```
res = make([]int, n); p = 0; q = 1
for x in a:
    if x > 0: res[p] = x; p += 2
    else:     res[q] = x; q += 2
```
- **Time:** O(n), **Space:** O(n) (the answer itself)

## Variant 2
Split into `pos` and `neg`. Alternate for `min(len(pos), len(neg))` pairs, then append the remaining ones.
**Time O(n), Space O(n)**.
