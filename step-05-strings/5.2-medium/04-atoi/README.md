# String to Integer (atoi)

**Problem (LeetCode 8):** Convert a string to a 32-bit signed integer, step by step:
1. Skip leading spaces.
2. Read an optional `+` or `−`.
3. Read digits until a non-digit or the end.
4. No digits → 0.
5. **Clamp** to `[−2³¹, 2³¹ − 1]` if it overflows.

| Input | Output |
|---|---|
| `"42"` | 42 |
| `"   -42"` | −42 |
| `"4193 with words"` | 4193 |
| `"words 987"` | 0 |
| `"-91283472332"` | −2147483648 (clamped) |

## Approach (iterative, optimal)
Follow the steps with an index `i`. Build `result = result*10 + digit`, and **check for overflow before it happens**: once `result` passes 2³¹, clamp and stop reading.
- **Time:** O(n), **Space:** O(1)

Recursive version (Step 7.1) uses the same rules, building the number one digit per call.
