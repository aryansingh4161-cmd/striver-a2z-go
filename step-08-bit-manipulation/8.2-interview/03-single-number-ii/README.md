# Single Number II (every other number appears three times)

**Problem (LeetCode 137):** `[2 2 3 2]` → 3

XOR doesn't work any more (three copies don't cancel).

## Approach 1: Hash map
Count; return the key with count 1. **Time O(n), Space O(n)**.

## Approach 2: Count each bit position (O(1) space)
For every bit position `b` (0..63), count how many numbers have that bit set. The numbers that appear three times add a multiple of 3 to that count. So `count % 3` is exactly the single number's bit.
- **Time:** O(64 × n), **Space:** O(1)

Works for negative numbers too because Go's `int` is 64-bit two's complement and we check all 64 bits.

## Approach 3: Sort
After sorting, triples sit together: check `a[i]` vs `a[i−1]` in steps of 3. **O(n log n)**.

## Approach 4: "ones and twos" buckets (optimal, tricky)
Keep two masks: `ones` = bits seen 1 time (mod 3), `twos` = bits seen 2 times.
```
ones = (ones ^ x) & ^twos
twos = (twos ^ x) & ^ones
```
After all numbers, `ones` is the answer.
- **Time:** O(n), **Space:** O(1)
