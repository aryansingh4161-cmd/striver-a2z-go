# Minimum Days to Make M Bouquets

**Problem (LeetCode 1482):** `bloomDay[i]` = the day flower `i` blooms. A bouquet needs `k` **adjacent** bloomed flowers. Return the minimum day you can make `m` bouquets, or −1 if impossible.

`bloomDay = [7 7 7 7 13 11 12 7]`, m = 2, k = 3 → 12

## Impossible case
If `m × k > n`, there aren't enough flowers → −1.

## Checker: `bouquets(day)`
Walk the array, counting consecutive bloomed flowers (`bloomDay[i] <= day`). Every time the count reaches `k`, make a bouquet and reset. A not-bloomed flower resets the count.

## Search space
Day `min(bloomDay)` .. `max(bloomDay)`. Later day → more flowers bloomed → once possible, always possible. ✅

- **Time:** O(n × log(max − min)), **Space:** O(1)

## Check day 12 on the example
Bloomed? `✓ ✓ ✓ ✓ ✗ ✓ ✓ ✓` → runs of 4 and 3 → 1 + 1 = 2 bouquets ✅. Day 11: `✓ ✓ ✓ ✓ ✗ ✓ ✗ ✓` → only 1 ❌. Answer **12**.
