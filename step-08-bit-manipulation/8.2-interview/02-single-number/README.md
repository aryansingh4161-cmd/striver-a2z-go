# Single Number (every other number appears twice)

Covered in detail in [Step 3](../../../step-03-arrays/3.1-easy/12-single-number/). Quick recap:

XOR everything. Pairs cancel (`x ^ x = 0`), the single number remains (`x ^ 0 = x`).
- **Time:** O(n), **Space:** O(1)

`[4 1 2 1 2]` → `4 ^ (1^1) ^ (2^2)` = **4**
