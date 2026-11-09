# Starting Point of the Loop

**Problem (LeetCode 142):** Return the node where the cycle begins, or nil.

## Approach 1: Hash set
The first node you see twice is the start. **Time O(n), Space O(n)**.

## Approach 2: Floyd's, phase two (optimal)
1. Run slow/fast until they meet (see [detect loop](../03-detect-loop/)).
2. Move `slow` back to `head`. Now move **both one step at a time**. Where they meet is the **start of the loop**.

**Why it works:**
Let `L1` = distance from head to the loop start, `d` = distance from the loop start to the meeting point, `C` = loop length.
When they meet: slow travelled `L1 + d`, fast travelled `2(L1 + d)`. Fast's extra distance is whole loops: `L1 + d = kC`. So `L1 = kC − d`: walking `L1` steps from the meeting point lands exactly on the loop start, the same as walking `L1` steps from the head.

- **Time:** O(n), **Space:** O(1)
