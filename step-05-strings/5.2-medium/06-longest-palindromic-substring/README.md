# Longest Palindromic Substring

**Problem (LeetCode 5):** Return the longest substring that is a palindrome.

`"babad"` → `"bab"` (or `"aba"`), `"cbbd"` → `"bb"`

## Approach 1: Check every substring
O(n²) substrings × O(n) check = **O(n³)**.

## Approach 2: Expand around centre (optimal for interviews)
Every palindrome has a **centre**: a single character (odd length) or a gap between two characters (even length). There are `2n − 1` centres.
From each centre, expand outward while both ends match.
```
for i in 0..n-1:
    expand(i, i)     // odd
    expand(i, i+1)   // even
```
- **Time:** O(n²), **Space:** O(1)

## Approach 3 (for later): Manacher's algorithm
O(n), but rarely expected in interviews.

## Dry run `"cbbd"`
Centre between the two b's: `s[1] == s[2]` ✓ → expand: `s[0]='c'` vs `s[3]='d'` ✗ → length 2 → `"bb"`.
