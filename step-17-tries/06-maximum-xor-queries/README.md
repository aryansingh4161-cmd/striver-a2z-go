# Maximum XOR With an Element From Array (Queries)

**Problem:** Array `nums` and queries `[x, m]`. For each query, return the max of `x XOR nums[j]` over all `nums[j] <= m`, or `-1` if no element is `<= m`.

```
nums = [0 1 2 3 4], queries = [[3 1] [1 3] [5 6]] → [3 3 7]
```

## Brute force
For each query scan the whole array: **O(q · n)**.

## Offline trick: sort, then grow the trie
The limit `m` changes per query, but a trie can't "hide" big numbers. So **answer the queries in increasing order of `m`**:
1. Sort `nums`.
2. Sort queries by `m` (remember each query's original index).
3. Keep a pointer into `nums`. For each query, insert every `nums[i] <= m` not inserted yet.
4. If the trie is empty → `-1`. Otherwise → the [bit-trie max XOR](../05-maximum-xor/) of `x`.
5. Put the answer back at the original index.

"Offline" = we're allowed to see all queries first and answer them in our own order.

## Dry run
Sorted nums `[0 1 2 3 4]`; queries by m: `(3,1)#0`, `(1,3)#1`, `(5,6)#2`
| query | inserted | best XOR |
|---|---|---|
| x=3, m=1 | 0, 1 | 3^0 = **3** |
| x=1, m=3 | + 2, 3 | 1^2 = **3** |
| x=5, m=6 | + 4 | 5^2 = **7** |
→ `[3 3 7]` ✅

- **Time:** O(n log n + q log q + (n + q) · 31), **Space:** O(n · 31)
