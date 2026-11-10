# Expression Add Operators

**Problem (LeetCode 282):** Insert `+`, `−` or `*` between the digits of `num` (or nothing, to join digits into longer numbers) so the expression equals `target`. Return all such expressions. Numbers can't have leading zeros (`"05"` is not allowed, `"0"` is fine).

`"123"`, 6 → `1+2+3`, `1*2*3`
`"105"`, 5 → `1*0+5`, `10-5`

## Approach: Backtracking that tracks the last term
At each position, choose how long the next number is (`num[i..j]`), then which operator goes before it.

The tricky part is `*`, because it binds tighter than `+`/`−`. Keep:
- `value` = result of the expression so far,
- `last` = the last **term** that was added (so a `*` can undo it).

| operator | new value | new last |
|---|---|---|
| `+ x` | `value + x` | `x` |
| `− x` | `value − x` | `−x` |
| `* x` | `value − last + last*x` | `last * x` |

Example: `1 + 2 * 3`. After `1+2`: value 3, last 2. Applying `*3`: `3 − 2 + 2*3 = 7` ✓.

**Leading zeros:** if `num[i] == '0'`, the number starting at `i` can only be `"0"` itself.
- **Time:** O(4ⁿ × n) (3 operators or "join" between each pair of digits), **Space:** O(n) depth
