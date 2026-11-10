# Merge K Sorted Lists

**Problem (LeetCode 23):** Merge `k` sorted linked lists into one sorted list.

`[1→4→5, 1→3→4, 2→6]` → `1→1→2→3→4→4→5→6`

Let N = total number of nodes.

## Approach 1: Collect, sort, rebuild
**O(N log N)**, O(N) extra space.

## Approach 2: Merge one by one
Merge list 1 with list 2, the result with list 3, … The growing result is walked again each time: **O(N × k)**.

## Approach 3: Min-heap of list heads (optimal)
Put the head of every list into a min-heap (by value). Repeatedly pop the smallest node, attach it to the result, and push that node's `Next` (if any).
- **Time:** O(N log k), **Space:** O(k)

(Divide and conquer, merging pairs like merge sort, also gives O(N log k).)

Uses the shared `list` package from [Step 6](../../../step-06-linked-list/list/).
