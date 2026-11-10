# Maximum XOR of Two Numbers in an Array

**Problem:** Non-negative integers `nums`. Return the maximum `nums[i] XOR nums[j]`.

```
[3 10 5 25 2 8] → 28 (5 XOR 25)
```

## Brute force
All pairs: **O(n²)**.

## Bit trie (binary trie)
XOR is big when the **high bits differ**. Store every number as a path of bits from the **most significant** (bit 30) down to bit 0. Each node has 2 children: bit 0 and bit 1.

For a number `x`, find its best partner greedily, bit by bit from the top:
- Want the **opposite** bit of `x` at this position (that makes this XOR bit 1).
- If that child exists, go there and set the bit in the answer.
- Otherwise, go the same way (this XOR bit stays 0).

Greedy from the top is correct because bit 30 is worth more than all lower bits together.

```
insert all numbers
best = 0
for x in nums:
    cur = root, val = 0
    for b = 30 down to 0:
        bit = (x >> b) & 1
        if cur.child[1−bit] != nil: val |= 1 << b; cur = cur.child[1−bit]
        else:                          cur = cur.child[bit]
    best = max(best, val)
```

## Dry run (5 vs the trie, 5 bits)
`5 = 00101`, `25 = 11001`. Walking for 5: wants 1 at bit 4 → 25 has it ✓; bit 3: wants 1 → 25 has 1 ✓; bit 2: 5 has 1, wants 0 → 25 has 0 ✓; bit 1: wants 1 → 25 has 0, take 0; bit 0: 5 has 1, wants 0 → 25 has 1, take 1.
XOR = `11100` = **28** ✅

- **Time:** O(n · 31), **Space:** O(n · 31) nodes
