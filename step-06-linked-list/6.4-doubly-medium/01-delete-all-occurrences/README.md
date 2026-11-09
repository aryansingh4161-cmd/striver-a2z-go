# Delete All Occurrences of a Key in a DLL

**Problem:** `10 ⇄ 4 ⇄ 10 ⇄ 6 ⇄ 10`, key 10 → `4 ⇄ 6`

## Approach (optimal)
Walk the list. When `cur.Val == key`:
- if `cur` is the head, move `head` to `cur.Next`,
- link `cur.Prev.Next = cur.Next` (if `Prev` exists),
- link `cur.Next.Prev = cur.Prev` (if `Next` exists),
- **save `cur.Next` before unlinking**, then continue from it.

- **Time:** O(n), **Space:** O(1)
