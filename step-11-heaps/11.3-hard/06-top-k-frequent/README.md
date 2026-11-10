# Top K Frequent Elements

**Problem (LeetCode 347):** Return the `k` most frequent elements (answer is unique).

`[1 1 1 2 2 3]`, k = 2 → `[1 2]`

## Approach 1: Count + sort by frequency
**O(n log n)**.

## Approach 2: Count + min-heap of size k
Push `(freq, value)` pairs; pop when the heap exceeds k. **O(n log k)**.

## Approach 3: Bucket sort (optimal)
A frequency is between 1 and n. Make `buckets[f]` = values that appear exactly `f` times. Walk `f` from n down to 1, collecting values until you have k.
- **Time:** O(n), **Space:** O(n)

Both the heap and bucket versions are in `solution.go`.
