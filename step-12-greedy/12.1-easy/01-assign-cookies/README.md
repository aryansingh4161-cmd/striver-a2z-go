# Assign Cookies

**Problem (LeetCode 455):** Child `i` is happy with a cookie of size ≥ `g[i]` (greed). Each child gets at most one cookie. Maximise happy children.

`g = [1 2 3]`, `s = [1 1]` → 1
`g = [1 2]`, `s = [1 2 3]` → 2

## Greedy rule
Sort both. Give the **smallest cookie that satisfies** the **least greedy** child. Using a bigger cookie than needed can only hurt later children.
```
sort g, sort s
i (children), j (cookies) = 0
while i < len(g) and j < len(s):
    if s[j] >= g[i]: i++     // child satisfied
    j++                      // cookie used or too small either way
return i
```
- **Time:** O(n log n + m log m), **Space:** O(1) besides sorting

## Why greedy is safe
If the smallest cookie can't satisfy the least greedy child, it can't satisfy anyone → throw it away. If it can, giving it to that child never blocks a better assignment.
