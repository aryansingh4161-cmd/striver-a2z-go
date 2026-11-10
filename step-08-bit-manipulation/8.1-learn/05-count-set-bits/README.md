# Count Set Bits

**Problem (LeetCode 191):** Count the 1s in the binary form of `n`. `13 = 1101` → 3.

## Approach 1: Check the last bit, shift right
```
while n > 0: count += n & 1; n >>= 1
```
**O(number of bits)** = O(log n).

## Approach 2: Brian Kernighan's trick (optimal)
`n & (n−1)` removes the lowest 1. Repeat until `n` is 0; count the steps.
```
while n != 0: n &= n - 1; count++
```
**O(number of set bits)**: faster when there are few 1s.

## Library
`bits.OnesCount(uint(n))` (package `math/bits`): uses a CPU instruction.

## Bonus: count set bits of every number from 0 to n (LeetCode 338)
`bits[i] = bits[i >> 1] + (i & 1)`: i has the same bits as i/2, plus its last bit. **O(n)** total.
