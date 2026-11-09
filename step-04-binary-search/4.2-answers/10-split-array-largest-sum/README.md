# Split Array: Largest Sum

**Problem (LeetCode 410):** Split the array into `k` non-empty contiguous subarrays so that the **largest subarray sum is as small as possible**. Return that sum.

`[7 2 5 10 8]`, k = 2 → 18 (`[7 2 5] [10 8]`)

## It's the same problem
This is [book allocation](../09-book-allocation/) with different words: books = elements, students = subarrays.
- Search space `max(a) .. sum(a)`.
- Checker: greedy count of parts with sum ≤ limit; need `parts <= k`.

- **Time:** O(n log sum), **Space:** O(1)

Recognising "this is the same pattern" is the real skill this section trains.
