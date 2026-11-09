# Clone a Linked List with Random Pointer

**Problem (LeetCode 138):** Each node has `Next` and `Random` (pointing to any node or nil). Make a **deep copy**: all new nodes, with the same structure, and no pointer into the old list.

## Approach 1: Hash map old → new
1. First pass: for every old node, create a copy; `m[old] = copy`.
2. Second pass: `m[old].Next = m[old.Next]`, `m[old].Random = m[old.Random]`.
- **Time:** O(n), **Space:** O(n)

## Approach 2: Interleave copies (optimal, O(1) extra space)
1. **Insert** each copy right after its original: `A → A' → B → B' → C → C'`.
2. **Random pointers:** `A'.Random = A.Random.Next` (the copy of A's random sits right after it).
3. **Separate** the two lists: restore the originals' `Next`, and link the copies together.

- **Time:** O(3n) = O(n), **Space:** O(1) (besides the new list itself)
