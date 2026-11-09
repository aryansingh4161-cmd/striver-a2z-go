# Count Frequency of Each Element

**Problem:** For every distinct element, return how many times it appears.

`[10 5 10 15 10 5]` → 10: 3, 5: 2, 15: 1

## Approach 1: Brute force with a visited array
For each element not yet visited, scan the rest of the array and count matches, marking them visited.
- **Time:** O(n²), **Space:** O(n)

## Approach 2: Hash map (optimal)
One pass: `freq[x]++`.
- **Time:** O(n) average, **Space:** O(k) for k distinct values

## Printing in order
Go maps iterate in random order. To print in **first-seen order**, also keep a slice of keys the first time you see each value. (Or sort the keys if you want sorted order.)
