# Step 11: Heaps (Priority Queues)

> Status: theory ready, problems coming ⏳

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
| # | Problem | Status |
|---|---|---|
| 1 | Implement a min-heap from scratch | ⏳ |
| 2 | Implement a max-heap | ⏳ |
| 3 | Check if an array represents a min-heap | ⏳ |
| 4 | Convert a min-heap to a max-heap | ⏳ |

### 11.2 Medium
| # | Problem | Status |
|---|---|---|
| 1 | Kth largest element in an array | ⏳ |
| 2 | Kth smallest element | ⏳ |
| 3 | Sort a K-sorted array | ⏳ |
| 4 | Merge K sorted lists | ⏳ |
| 5 | Replace each element by its rank | ⏳ |
| 6 | Task scheduler | ⏳ |
| 7 | Hand of straights | ⏳ |

### 11.3 Hard
| # | Problem | Status |
|---|---|---|
| 1 | Design Twitter | ⏳ |
| 2 | Minimum cost to connect ropes | ⏳ |
| 3 | Kth largest element in a stream | ⏳ |
| 4 | Maximum sum combinations | ⏳ |
| 5 | Find median from data stream | ⏳ |
| 6 | Top K frequent elements | ⏳ |
