# Partition Array for Maximum Sum

**Problem:** Split `arr` into contiguous blocks, each of length **at most `k`**. Then every element of a block becomes that block's **maximum**. Maximise the total sum after this.

```
arr = [1 15 7 9 2 5 10], k = 3 → 84
blocks [1 15 7] [9] [2 5 10] → [15 15 15 9 10 10 10] = 84
```

## Front partition again
Same shape as [Palindrome Partitioning II](../05-palindrome-partitioning-ii/): choose the **first** block, then recurse on the rest.
```
f(n) = 0
f(i) = max over len = 1..k (with i + len ≤ n) of
       maxOf(arr[i .. i+len−1]) · len + f(i + len)
answer = f(0)
```
Keep a running max while `len` grows, so each choice costs O(1).

## Stages
- Recursion: O(kⁿ)
- Memo / table: n states × k choices → **O(n · k)**, Space O(n)

## Dry run (tail of the example)
- f(7) = 0
- f(6) = 10·1 = 10
- f(5) = max(5 + f(6) = 15, 10·2 + f(7) = 20) = 20
- f(4) = max(2 + 20, 5·2 + 10, 10·3 + 0) = 30
- ... f(0) = **84** ✅
