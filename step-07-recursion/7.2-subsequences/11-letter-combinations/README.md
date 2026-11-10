# Letter Combinations of a Phone Number

**Problem (LeetCode 17):** Digits 2–9 map to letters like an old phone keypad. Return every possible letter string.

| 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
|---|---|---|---|---|---|---|---|
| abc | def | ghi | jkl | mno | pqrs | tuv | wxyz |

`"23"` → `ad ae af bd be bf cd ce cf`

## Approach: One recursion level per digit
```
f(i):
    if i == len(digits): save current string; return
    for each letter of digits[i]:
        add letter; f(i+1); remove letter
```
Empty input → empty answer (not `[""]`).
- **Time:** O(4ⁿ × n) (up to 4 letters per digit), **Space:** O(n) depth
