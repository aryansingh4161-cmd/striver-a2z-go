# Check Odd or Even

The last bit (bit 0) is worth 1; every other bit is worth an even number. So:
**n is odd ⇔ its last bit is 1 ⇔ `n & 1 == 1`.**

Works for negative numbers in Go too (two's complement keeps the last bit). `n % 2 == 1` would **fail** for negative odd numbers in Go, because `-3 % 2 == -1`.
- **Time:** O(1)
