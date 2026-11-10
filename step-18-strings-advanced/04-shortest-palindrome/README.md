# Shortest Palindrome (Minimum Characters to Add in Front)

**Problem:** Add characters **only at the front** of `s` to make it a palindrome. Return the shortest such palindrome (Striver's version asks for the **count** of added characters; both are here).

```
"aacecaaa" → "aaacecaaa" (add 1 char)
"abcd"     → "dcbabcd"   (add 3 chars)
```

## Key insight
Whatever we add goes in front, so the original string must end up as the right part. The part of `s` we **don't** need to mirror is its **longest palindromic prefix**.
- `s = P + rest`, where `P` is the longest prefix that's a palindrome.
- Answer = `reverse(rest) + s`. Characters added = `len(rest) = n − len(P)`.

## Brute force
Try prefix lengths from `n` down to 1, check if it's a palindrome: **O(n²)**.

## KMP trick: O(n)
Build `t = s + "#" + reverse(s)`. Then [LPS](../03-kmp-lps/)`[last]` of `t` = longest prefix of `s` that is also a suffix of `reverse(s)`.
A prefix of `s` that equals a suffix of `reverse(s)` reads the same reversed: it's a **palindromic prefix**. And it's the **longest** one.

The `#` stops the match from running across the middle (the prefix-suffix can't be longer than `s`).

```
L = LPS(s + "#" + reverse(s))[last]
added = n − L
answer = reverse(s[L:]) + s
```

## Dry run (`abcd`)
`t = "abcd#dcba"`, last LPS = 1 ("a") → longest palindromic prefix "a" → add `reverse("bcd") = "dcb"` → **"dcbabcd"** (3 added) ✅

- **Time:** O(n), **Space:** O(n)
