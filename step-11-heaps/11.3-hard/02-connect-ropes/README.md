# Minimum Cost to Connect Ropes

**Problem:** Joining two ropes of lengths `a` and `b` costs `a + b` and makes a rope of length `a + b`. Join all ropes into one with minimum total cost.

`[4 3 2 6]` → 29
(2+3 = 5; 4+5 = 9; 6+9 = 15; total 5 + 9 + 15 = 29)

## Approach: Always join the two shortest (greedy + min-heap)
A rope's length is paid again every time its rope is joined later, so long ropes should be joined **as late as possible**. Always combining the two smallest achieves that (same idea as **Huffman coding**).
```
push all lengths into a min-heap
while size > 1:
    a = pop, b = pop
    cost += a + b
    push a + b
```
- **Time:** O(n log n), **Space:** O(n)
