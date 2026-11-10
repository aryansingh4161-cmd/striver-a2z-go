# KMP Algorithm (LPS Array)

**Problem:** Find all occurrences of `p` in `s` in **O(n + m)**, by never moving backwards in the text.

## The LPS array
`lps[i]` = length of the longest **proper prefix** of `p[0..i]` that is also a **suffix** of `p[0..i]` ("proper" = not the whole string).

```
p   = a a b a a a
lps = 0 1 0 1 2 2
          ↑ lps[4] = 2: "aabaa" starts and ends with "aa"
```

## Why it helps
When matching `s` against `p`, suppose `j` characters of `p` matched and then a mismatch happens. Those `j` characters are `p[0..j−1]`, and they're also the **last** `j` characters read from `s`. The longest prefix of `p` that we can still be in the middle of is `lps[j−1]`: so jump `j = lps[j−1]` and **keep `i` where it is**.

## Building LPS (the pattern matched against itself)
```
len = 0               // length of the current prefix-suffix
for i = 1..m−1:
    while len > 0 and p[i] != p[len]: len = lps[len−1]   // fall back
    if p[i] == p[len]: len++
    lps[i] = len
```

## Searching
```
j = 0                 // characters of p matched so far
for i = 0..n−1:
    while j > 0 and s[i] != p[j]: j = lps[j−1]
    if s[i] == p[j]: j++
    if j == m: found at i − m + 1; j = lps[j−1]          // keep going for overlaps
```

## Dry run (LPS of `aabaaa`)
| i | p[i] | len before | action | lps[i] |
|---|---|---|---|---|
| 1 | a | 0 | p[1] = p[0] → 1 | 1 |
| 2 | b | 1 | b ≠ p[1]=a → len = lps[0] = 0; b ≠ a | 0 |
| 3 | a | 0 | match → 1 | 1 |
| 4 | a | 1 | match → 2 | 2 |
| 5 | a | 2 | a ≠ p[2]=b → len = lps[1] = 1; a = p[1] → 2 | 2 |

## Complexity
Each step either moves `i` forward or makes `len`/`j` smaller; `len`/`j` can't drop more times than they grew. → **O(n + m)** time, **O(m)** space.

## KMP vs Z vs Rabin-Karp
| | Time | Notes |
|---|---|---|
| Rabin-Karp | O(n + m) expected | hashing, verify on match |
| Z-function | O(n + m) | build on `p + "$" + s` |
| KMP | O(n + m) | only needs `lps` of the pattern |
