# Intersection Point of Two Linked Lists

**Problem (LeetCode 160):** Two lists may merge into one shared tail (Y shape). Return the first shared **node**, or nil.

```
A:  4 → 1 ↘
            8 → 4 → 5
B: 5 → 6 → 1 ↗
```
Answer: the node with value 8 (compare **nodes**, not values).

## Approach 1: Hash set
Put all nodes of A in a set; walk B and return the first node found in the set. **Time O(n + m), Space O(n)**.

## Approach 2: Length difference
Find both lengths, move the longer list's pointer ahead by the difference, then move both together until they meet. **Time O(n + m), Space O(1)**.

## Approach 3: Switch heads (optimal, shortest code)
Walk `a` and `b` together. When `a` reaches the end, jump to B's head; when `b` reaches the end, jump to A's head.
Both pointers travel `lenA + lenB` steps in total, so they line up and meet at the intersection, or both become nil at the same time if there's none.
- **Time:** O(n + m), **Space:** O(1)
