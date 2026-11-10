# LRU Cache

**Problem (LeetCode 146):** A cache with fixed `capacity`:
- `Get(key)`: return the value, or −1. Using a key makes it the **most recently used**.
- `Put(key, value)`: insert or update. If the cache is over capacity, evict the **least recently used** key.
Both must be **O(1)**.

## Design: Hash map + doubly linked list
- The **list** keeps usage order: most recent right after `head`, least recent right before `tail`. (`head` and `tail` are dummy nodes, so there are no nil checks.)
- The **map** goes from `key → node` so we can jump to any node in O(1).

| Operation | Steps |
|---|---|
| Get(key) | not in map → −1. Otherwise **move the node to the front**, return its value. |
| Put(key, v) | in map → update value, move to front. Else: if full, **remove `tail.prev`** (the LRU) from list and map; then insert a new node at the front. |

A doubly linked list is needed so a node can be **unlinked in O(1)** from anywhere (we know its `prev`).

- **Time:** O(1) per operation, **Space:** O(capacity)

## Example (capacity 2)
`put(1,1) put(2,2) get(1)=1 put(3,3)` → evicts 2 (1 was used more recently) → `get(2) = −1`.
