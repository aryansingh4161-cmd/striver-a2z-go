# Kth Largest Element in an Array

**Problem (LeetCode 215):** `[3 2 1 5 6 4]`, k = 2 → 5

## Approach 1: Sort
Sort descending, take index `k−1`. **O(n log n)**.

## Approach 2: Min-heap of size k (common interview answer)
Keep the **k largest seen so far** in a min-heap. Push each element; if the heap grows past `k`, pop the smallest. At the end the top is the k-th largest.
- **Time:** O(n log k), **Space:** O(k)

It feels backwards (a **min**-heap for the **largest**) but the heap's top is the weakest of the top-k: exactly the k-th largest.

## Approach 3: Quickselect (average O(n))
Partition like [quick sort](../../../step-02-sorting/07-quick-sort/), but recurse only into the side that contains position `n−k`. Average **O(n)**, worst O(n²) (random pivot makes that unlikely).
