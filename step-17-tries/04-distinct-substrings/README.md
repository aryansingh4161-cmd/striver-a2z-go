# Number of Distinct Substrings

**Problem:** Count the distinct **non-empty** substrings of `s`.

```
"abab" → 7 (a, b, ab, ba, aba, bab, abab)
"aaa"  → 3 (a, aa, aaa)
```
(Some versions, like Striver's GFG problem, also count the **empty** string, so their answer is ours + 1.)

## Brute force
Put every substring `s[i..j]` into a hash set. There are O(n²) substrings, each costs O(n) to build/hash → **O(n³)** time, O(n³) memory in the worst case.

## Trie of all substrings
Every substring is a **prefix of some suffix**. So insert each suffix `s[i..]` into a trie, walking one letter at a time:
- Every time we **create** a new node, we've found a substring never seen before → `count++`.
- Existing nodes are substrings already counted.

```
count = 0
for i in 0..n−1:
    cur = root
    for j in i..n−1:
        if cur.child[s[j]] is nil: create it; count++
        cur = cur.child[s[j]]
```
**Answer = number of nodes created** (excluding the root).

## Dry run (`"abab"`)
| suffix | new nodes | count |
|---|---|---|
| abab | a, ab, aba, abab | 4 |
| bab | b, ba, bab | 7 |
| ab | – (already there) | 7 |
| b | – | 7 |
→ **7** ✅

- **Time:** O(n²), **Space:** O(n²) nodes in the worst case

(For n up to 10⁵ you'd use a suffix array / suffix automaton, beyond this sheet.)
