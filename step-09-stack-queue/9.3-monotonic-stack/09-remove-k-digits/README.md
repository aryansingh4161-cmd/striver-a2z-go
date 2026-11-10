# Remove K Digits

**Problem (LeetCode 402):** Remove exactly `k` digits from the number string to make it as **small** as possible. No leading zeros in the answer; empty → `"0"`.

`"1432219"`, k = 3 → `"1219"`
`"10200"`, k = 1 → `"200"`

## Key insight
A number is smaller when its **leftmost** digits are smaller. So whenever a digit is **bigger than the one after it**, removing it helps the most.

## Approach: Monotonic increasing stack (greedy, optimal)
For each digit `d`:
- while `k > 0` and the top of the stack is **bigger** than `d`: pop it, `k--`
- push `d`

Afterwards:
- if `k` is still > 0 (the digits were increasing), remove from the **end**,
- strip leading zeros,
- empty → `"0"`.

- **Time:** O(n), **Space:** O(n)

## Dry run `"1432219"`, k = 3
| d | action | stack | k |
|---|---|---|---|
| 1 | push | 1 | 3 |
| 4 | push | 14 | 3 |
| 3 | pop 4 | 13 | 2 |
| 2 | pop 3 | 12 | 1 |
| 2 | push | 122 | 1 |
| 1 | pop 2 | 121 | 0 |
| 9 | push | 1219 | 0 |
