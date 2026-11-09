# Nth Root of an Integer

**Problem:** Return the integer `x` with `xⁿ = m`, or −1 if there's no integer root.

| n, m | answer |
|---|---|
| 3, 27 | 3 |
| 4, 69 | −1 |

## Approach: Binary search on x in 1..m
Compare `midⁿ` with `m`:
- equal → found
- smaller → go right
- bigger → go left

**Overflow trick:** don't compute `midⁿ` fully. Multiply step by step and **stop as soon as the product exceeds m** (return "too big"). This keeps numbers small.
- **Time:** O(n log m), **Space:** O(1)
