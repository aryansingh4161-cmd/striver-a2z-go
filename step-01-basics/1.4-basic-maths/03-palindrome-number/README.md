# Check Palindrome Number

**Problem:** Return `true` if `n` reads the same forwards and backwards.

| Input | Output |
|---|---|
| 121 | true |
| 123 | false |
| -121 | false (the `-` is only at the front) |

## Intuition
A palindrome equals its own reverse. Reuse the [reverse number](../02-reverse-number/) logic.

## Approach
1. If `n < 0` → false.
2. Keep a copy `original := n`.
3. Reverse `n`.
4. Return `rev == original`.

- **Time:** O(log₁₀ n)
- **Space:** O(1)

## Dry run (n = 121)
Reverse: 1 → 12 → 121. `121 == 121` → **true**.

## Common mistake
Reversing `n` destroys it (it becomes 0). **Save a copy first**, then compare with the copy.
