# Step 11: Heaps (Priority Queues)

> Status: ✅ theory and all 17 problems done

## What it is
A **complete binary tree** stored in a slice where every parent is ≤ its children (**min-heap**) or ≥ (**max-heap**). The smallest (or largest) element is always at index 0.

For index `i`: left child `2i+1`, right child `2i+2`, parent `(i−1)/2`.

| Operation | Cost |
|---|---|
| peek min/max | O(1) |
| push | O(log n) (bubble up) |
| pop min/max | O(log n) (move last to top, sift down) |
| build heap from n items | O(n) |

## In Go: `container/heap`
You implement 5 methods once; the package does the heap work. Template in [docs/go-for-dsa.md](../docs/go-for-dsa.md#8-heap-priority-queue).
- Max-heap: flip `Less` to `h[i] > h[j]`.
- Heap of pairs/structs: store structs and compare on the field you need.

## When to think "heap"
- "K largest / K smallest / K most frequent" → keep a heap of size **K**: O(n log K).
  - K **largest** → use a **min**-heap of size K; pop when it grows past K. The top is the K-th largest.
- Repeatedly take the smallest/largest (merge K sorted lists, connect ropes, task scheduler).
- Running median → **two heaps**: a max-heap for the lower half and a min-heap for the upper half.

## Problems

### 11.1 Learning
| # | Problem | Key idea |
|---|---|---|
| 1 | [Min-heap from scratch](11.1-learning/01-min-heap-from-scratch/) | sift up / sift down, O(n) build |
| 2 | [Max-heap](11.1-learning/02-max-heap/) | flip the comparisons / Less |
| 3 | [Check if an array is a min-heap](11.1-learning/03-check-min-heap/) | check non-leaves only |
| 4 | [Convert min-heap to max-heap](11.1-learning/04-min-to-max-heap/) | just build a max-heap, O(n) |

### 11.2 Medium
| # | Problem | Key idea |
|---|---|---|
| 1 | [Kth largest element](11.2-medium/01-kth-largest/) | min-heap of size k / quickselect |
| 2 | [Kth smallest element](11.2-medium/02-kth-smallest/) | max-heap of size k |
| 3 | [Sort a K-sorted array](11.2-medium/03-sort-k-sorted-array/) | window heap of size k+1 |
| 4 | [Merge K sorted lists](11.2-medium/04-merge-k-sorted-lists/) | heap of list heads |
| 5 | [Replace each element by its rank](11.2-medium/05-replace-by-rank/) | sort + map |
| 6 | [Task scheduler](11.2-medium/06-task-scheduler/) | (maxFreq−1)(n+1) + countMax |
| 7 | [Hand of straights](11.2-medium/07-hand-of-straights/) | smallest card starts a group |

### 11.3 Hard
| # | Problem | Key idea |
|---|---|---|
| 1 | [Design Twitter](11.3-hard/01-design-twitter/) | merge followees' tweets with a max-heap |
| 2 | [Minimum cost to connect ropes](11.3-hard/02-connect-ropes/) | always join the two shortest |
| 3 | [Kth largest in a stream](11.3-hard/03-kth-largest-stream/) | persistent min-heap of size k |
| 4 | [Maximum sum combinations](11.3-hard/04-max-sum-combinations/) | grid frontier with a max-heap |
| 5 | [Median from data stream](11.3-hard/05-median-data-stream/) | two heaps: low max-heap, high min-heap |
| 6 | [Top K frequent elements](11.3-hard/06-top-k-frequent/) | heap of size k / buckets by frequency |
