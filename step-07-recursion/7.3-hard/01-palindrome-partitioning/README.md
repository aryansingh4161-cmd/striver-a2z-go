# Palindrome Partitioning

**Problem (LeetCode 131):** Split `s` into pieces so that **every piece is a palindrome**. Return all such splits.

`"aab"` → `[a a b] [aa b]`

## Approach: Try every first cut
At position `start`, try every end `i` (from `start` to `n−1`). If `s[start..i]` is a palindrome, take it as the next piece and recurse from `i+1`. When `start == n`, the whole string is split → save.
```
f(start):
    if start == n: save path; return
    for i = start .. n-1:
        if isPalindrome(start, i):
            path.push(s[start..i]); f(i+1); path.pop()
```
- **Time:** O(2ⁿ × n) (every cut position can be on/off; each check costs O(n)), **Space:** O(n) depth

## Recursion tree "aab"
```
f(0)
├─ "a"  → f(1)
│        ├─ "a" → f(2) → "b" → f(3) ✓  [a a b]
│        └─ "ab" ✗
├─ "aa" → f(2) → "b" → f(3) ✓  [aa b]
└─ "aab" ✗
```
(Step 16 has "Palindrome Partitioning II": the **minimum** number of cuts, solved with DP.)
