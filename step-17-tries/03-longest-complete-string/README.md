# Longest Word with All Prefixes Present (Complete String)

**Problem:** Given a list of words, find the **longest** word such that **every prefix** of it is also in the list. If several have the same length, return the **lexicographically smallest**. If none, return `"None"`.

```
["n" "ni" "nin" "ninj" "ninja" "ninga"] → "ninja"
   ("ninga" fails: "ning" is missing)
["ab" "bc"] → "None"  ("a" and "b" are missing)
```
(LeetCode "Longest Word in Dictionary" is the same idea, returning `""` instead of `"None"`.)

## Brute force
For each word, check all its prefixes in a hash set: O(N · L²) (building each prefix costs O(L)).

## Trie
1. Insert every word.
2. For each word, walk its letters; **every** node on the path must have `end = true`. One walk = O(L), so total **O(N · L)**.
3. Keep the best by (longer, then smaller).

```
for w in words:
    cur = root; ok = true
    for c in w:
        cur = cur.child[c]
        if !cur.end: ok = false; break
    if ok and (len(w) > len(best) or (len(w) == len(best) and w < best)): best = w
```

## Dry run
- "ninja": n ✓, ni ✓, nin ✓, ninj ✓, ninja ✓ → candidate (5)
- "ninga": n, ni, nin ✓, ning ✗ → rejected
→ **"ninja"** ✅

- **Time:** O(N · L), **Space:** O(N · L · 26)
