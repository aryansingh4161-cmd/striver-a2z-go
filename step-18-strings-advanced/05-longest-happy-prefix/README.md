# Longest Happy Prefix

**Problem:** A "happy prefix" is a non-empty prefix that is **also a suffix** (and isn't the whole string). Return the longest one, or `""`.

```
"level"   → "l"
"ababab"  → "abab"
"abc"     → ""
```

## Brute force
Try lengths `n−1` down to 1; compare `s[:l]` with `s[n−l:]`: **O(n²)**.

## It's literally the last LPS value
[LPS](../03-kmp-lps/)`[n−1]` is defined as the longest **proper** prefix of the whole string that's also a suffix. That's exactly the happy prefix:
```
answer = s[ : LPS(s)[n−1] ]
```

## Dry run (`ababab`)
LPS = `0 0 1 2 3 4` → last = 4 → **"abab"** ✅ (overlapping is fine: prefix "abab" and suffix "abab" share the middle "ab")

- **Time:** O(n), **Space:** O(n)

(Z-function also works: the answer is the largest `z[i]` with `i + z[i] == n`. Rabin-Karp works too: compare prefix and suffix hashes for each length.)
