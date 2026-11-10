# Min Stack

**Problem (LeetCode 155):** A stack that also returns the **minimum** element in O(1).

## Approach 1: Store pairs (value, minSoFar)
Each entry remembers the minimum of everything at or below it.
- push(x): `min = x` if empty, else `min(x, top.min)`; push `(x, min)`
- getMin(): `top.min`
- **All O(1), Space O(2n)**

## Approach 2: One value per entry with an encoding trick (O(n) space, no pairs)
Keep a variable `mini`. When pushing `x < mini`, push the **encoded** value `2x − mini` (which is always less than `x`) and set `mini = x`. When popping a value smaller than `mini`, it's encoded: restore the previous minimum with `mini = 2·mini − value`.
This is a classic interview follow-up, but the pairs version is clearer and what you should write first.

## Dry run (pairs): push 5, 3, 7, 2
| stack (bottom → top) | getMin |
|---|---|
| (5,5) | 5 |
| (5,5) (3,3) | 3 |
| (5,5) (3,3) (7,3) | 3 |
| … (2,2) | 2 |
pop → back to 3 automatically.
