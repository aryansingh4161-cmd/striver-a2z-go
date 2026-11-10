# Minimum Insertions / Deletions to Convert A to B

**Problem:** Turn string `a` into string `b` using only **delete** and **insert** operations. Minimum number of operations? (LeetCode 583 "Delete Operation for Two Strings" is the same count.)

```
a = "abcd", b = "anc" → 3 (delete b, delete d, insert n)
```

## Idea
Characters in the [LCS](../01-lcs/) can stay where they are. Everything else has to change:
- Delete every char of `a` not in the LCS: `n − LCS`
- Insert every char of `b` not in the LCS: `m − LCS`

```
answer = (n − LCS) + (m − LCS) = n + m − 2·LCS
```

## Dry run
`abcd` / `anc`: LCS = "ac" (2) → 4 + 3 − 4 = **3** ✅

- **Time:** O(n · m), **Space:** O(m)
