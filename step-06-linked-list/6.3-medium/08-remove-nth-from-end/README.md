# Remove N-th Node from the End

**Problem (LeetCode 19):** `1 → 2 → 3 → 4 → 5`, n = 2 → `1 → 2 → 3 → 5`

## Approach 1: Count, then delete
Length `L`; the node to delete is at position `L − n + 1` from the start. Walk to the one before it. **Two passes**.

## Approach 2: Gap of n with two pointers (optimal, one pass)
1. Use a **dummy** node before head (handles deleting the head).
2. Move `fast` n steps ahead of `slow`.
3. Move both until `fast.Next == nil`. Now `slow` is right **before** the target.
4. `slow.Next = slow.Next.Next`.

- **Time:** O(n), **Space:** O(1)

## Dry run n = 2
| step | slow | fast |
|---|---|---|
| after gap | dummy | 2 |
| move | 1 | 3 |
| move | 2 | 4 |
| move | 3 | 5 (last) → stop |
Delete `3.Next` (= 4) → `1 2 3 5`.
