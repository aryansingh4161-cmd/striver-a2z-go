# LFU Cache

**Problem (LeetCode 460):** Like [LRU](../04-lru-cache/), but evict the **least frequently used** key. Ties (same frequency) → evict the **least recently used** among them. `Get` and `Put` must be O(1).

## Design
- `items`: `key → node` (node holds key, value, freq)
- `lists`: `freq → doubly linked list` of nodes with that frequency, most recent at the front
- `minFreq`: the smallest frequency currently in the cache

| Operation | Steps |
|---|---|
| **touch(node)** (on every Get/Put hit) | remove it from `lists[freq]`; if that list is now empty **and** `freq == minFreq`, then `minFreq++`. `freq++`, add to the front of `lists[freq]`. |
| Get(key) | missing → −1; else touch, return value |
| Put(key, v) | exists → update value, touch. Else: if full, evict the **back** of `lists[minFreq]` (LFU, and LRU among ties). Insert the new node with `freq = 1` and set **`minFreq = 1`**. |

Why is `minFreq` easy to maintain? It only changes in two ways: a new key resets it to 1, and touching the last node of the min list raises it by exactly 1.

- **Time:** O(1) per operation, **Space:** O(capacity)
