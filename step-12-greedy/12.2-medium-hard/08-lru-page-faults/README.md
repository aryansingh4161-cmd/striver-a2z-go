# Page Faults in LRU (Least Recently Used)

**Problem:** Memory holds `capacity` pages. Process page requests in order. If a page isn't in memory, it's a **page fault**: load it, and if memory is full, evict the **least recently used** page. Count page faults.

pages `[5 0 1 3 2 4 1 0 5]`, capacity 4 → 8

## Approach 1: Slice ordered by recency
Keep a slice with the most recent page at the end.
- hit → move that page to the end (O(capacity))
- miss → fault; if full remove index 0 (the LRU); append
- **Time:** O(n × capacity), **Space:** O(capacity)

## Approach 2: Hash map + doubly linked list
Exactly the [LRU cache](../../../step-09-stack-queue/9.4-implementation/04-lru-cache/) from Step 9: O(1) per request.
- **Time:** O(n), **Space:** O(capacity)

Greedy idea: evicting the page whose last use is furthest in the past.
