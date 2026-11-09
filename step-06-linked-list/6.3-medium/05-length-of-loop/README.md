# Length of the Loop

**Problem:** Return the number of nodes in the cycle, or 0 if there's none.

## Approach 1: Hash map of node → step number
Walk with a step counter, storing when each node was first seen. When you reach a seen node: `length = step − firstSeen[node]`.
**Time O(n), Space O(n)**.

## Approach 2: Floyd's + one lap (optimal)
1. Find the meeting point with slow/fast.
2. From there, walk one pointer around the loop counting steps until it returns to the meeting point.
- **Time:** O(n), **Space:** O(1)
