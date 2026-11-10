# Minimum Insertions to Make a String Palindrome

**Problem:** Fewest characters to insert (anywhere) so `s` becomes a palindrome.

```
"abcaa" → 2 ("abcaa" → "aabcbaa"... e.g. insert b and a)
"mbadm" → 2 ("mbdadbm")
"zzazz" → 0
```

## Idea
Keep the part that's **already** a palindrome and only fix the rest.
- The [longest palindromic subsequence](../04-longest-palindromic-subsequence/) is the biggest part we can keep untouched.
- Every other character needs a "mirror partner" inserted on the other side: one insertion each.

```
answer = n − LPS(s)
```

## Dry run
`mbadm`: LPS = "mam" or "mdm" (3) → 5 − 3 = **2** ✅

- **Time:** O(n²), **Space:** O(n)

(The same answer = minimum **deletions** to make it a palindrome: delete the characters not in the LPS.)
