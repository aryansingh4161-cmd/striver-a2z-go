# Check if the i-th Bit is Set (plus set / clear / toggle)

Bits are numbered from the right, starting at 0. `13 = 1101`: bits 0, 2, 3 are set.

## Check
- **Left shift a mask:** `n & (1 << i) != 0`
- **Right shift n:** `(n >> i) & 1 == 1`

Both are **O(1)**.

## Related one-liners
| Task | Go |
|---|---|
| set bit i | `n \| (1 << i)` |
| clear bit i | `n &^ (1 << i)` (Go's AND-NOT) or `n & ^(1 << i)` |
| toggle bit i | `n ^ (1 << i)` |

## Dry run n = 13 (1101), i = 2
`1 << 2 = 0100`; `1101 & 0100 = 0100` ≠ 0 → **set**.
