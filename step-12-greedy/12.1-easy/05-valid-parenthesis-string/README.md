# Valid Parenthesis String (with `*`)

**Problem (LeetCode 678):** String of `(`, `)` and `*`. A `*` can be `(`, `)` or empty. Can the string be valid?

`"(*)"` → true, `"(*))"` → true, `"(()"` → false

## Approach 1: Recursion
At each `*` try all three choices, tracking the open count. **O(3ⁿ)**. (With memoisation on (index, count) it becomes O(n²): a DP.)

## Approach 2: Track a range of possible open counts (greedy, optimal)
Keep `lo` and `hi` = the smallest and largest number of unmatched `(` possible so far.
- `(` → `lo++`, `hi++`
- `)` → `lo--`, `hi--`
- `*` → `lo--` (treat as `)`), `hi++` (treat as `(`)
- If `hi < 0` → too many `)` no matter what → false.
- Clamp `lo` at 0 (a negative count isn't a real option; those choices just aren't used).
At the end, valid if `lo == 0` (zero open is reachable).
- **Time:** O(n), **Space:** O(1)

## Dry run `"(*))"`
| ch | lo | hi |
|---|---|---|
| ( | 1 | 1 |
| * | 0 | 2 |
| ) | 0 (clamped) | 1 |
| ) | 0 (clamped) | 0 |
lo == 0 → **true**
