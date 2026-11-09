# Check Palindrome String (recursion)

**Problem:** Check if a string reads the same forwards and backwards.

| Input | Output |
|---|---|
| "madam" | true |
| "hello" | false |

**LeetCode 125 version:** ignore everything except letters and digits, and ignore case. `"A man, a plan, a canal: Panama"` → true.

## Approach
Compare the first and last characters, then recurse on the inside.
```
f(i):
    if i >= n/2: return true
    if s[i] != s[n-1-i]: return false
    return f(i + 1)
```
- **Time:** O(n), **Space:** O(n) stack

## For LeetCode 125
First build a cleaned string: keep only letters/digits, lowercase them. Then run the same check.
