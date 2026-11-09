# Roman to Integer (and Integer to Roman)

| Symbol | I | V | X | L | C | D | M |
|---|---|---|---|---|---|---|---|
| Value | 1 | 5 | 10 | 50 | 100 | 500 | 1000 |

Normally symbols go from largest to smallest and you add them. **Subtraction rule:** a smaller symbol **before** a larger one is subtracted: IV = 4, IX = 9, XL = 40, XC = 90, CD = 400, CM = 900.

## Roman → integer (LeetCode 13)
`"MCMXCIV"` → 1994

Walk left to right. If the current value is **smaller than the next** one, subtract it; otherwise add it.
- **Time:** O(n), **Space:** O(1)

| ch | value | next | action | total |
|---|---|---|---|---|
| M | 1000 | C | add | 1000 |
| C | 100 | M | subtract | 900 |
| M | 1000 | X | add | 1900 |
| X | 10 | C | subtract | 1890 |
| C | 100 | I | add | 1990 |
| I | 1 | V | subtract | 1989 |
| V | 5 | – | add | **1994** |

## Integer → Roman (LeetCode 12)
Greedy with a table that **includes the subtraction pairs**:
`1000 M, 900 CM, 500 D, 400 CD, 100 C, 90 XC, 50 L, 40 XL, 10 X, 9 IX, 5 V, 4 IV, 1 I`
Repeatedly take the largest value that fits.
- **Time:** O(1) (numbers ≤ 3999), **Space:** O(1)
