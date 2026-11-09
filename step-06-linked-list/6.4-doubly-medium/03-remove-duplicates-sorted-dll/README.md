# Remove Duplicates from a Sorted DLL

**Problem:** `1 ⇄ 1 ⇄ 1 ⇄ 2 ⇄ 3 ⇄ 3 ⇄ 4` → `1 ⇄ 2 ⇄ 3 ⇄ 4`

## Approach (optimal)
For each node `cur`, skip ahead with `nxt` while `nxt.Val == cur.Val`. Then link `cur.Next = nxt` and, if `nxt` isn't nil, `nxt.Prev = cur`. Move `cur = nxt`.
- **Time:** O(n), **Space:** O(1)
