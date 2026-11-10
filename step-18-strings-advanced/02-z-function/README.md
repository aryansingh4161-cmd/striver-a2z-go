# Z-Function

**Problem:** For a string `s`, build `z` where `z[i]` = length of the longest substring starting at `i` that is also a **prefix** of `s`. (`z[0]` is usually set to 0 or n.) Then use it to find a pattern.

```
s = "aabxaab"
z = [0 1 0 0 3 1 0]
      ↑ z[4] = 3: "aab" at index 4 matches the prefix "aab"
```

## Brute force
For each `i`, extend while `s[i+k] == s[k]`: **O(n²)**.

## The Z-box trick: O(n)
Keep the window `[l, r)` = the rightmost segment found so far that matches a prefix (`s[l..r−1] == s[0..r−l−1]`).

For a new `i`:
- **If `i < r`** (inside the box): `s[i..r−1]` equals `s[i−l..r−l−1]`, whose answer we already know. So start with `z[i] = min(r − i, z[i − l])`.
- Then **extend** naively while characters match.
- If the match goes past `r`, move the box: `l = i`, `r = i + z[i]`.

Every successful extension pushes `r` to the right, and `r` only goes up to n → **O(n)** total.

```
l, r = 0, 0
for i = 1..n−1:
    if i < r: z[i] = min(r − i, z[i − l])
    while i + z[i] < n and s[z[i]] == s[i + z[i]]: z[i]++
    if i + z[i] > r: l, r = i, i + z[i]
```

## Pattern search with Z
Build `z` of `p + "$" + s` (`$` is a character in neither string). Wherever `z[i] == m`, the pattern starts at `i − m − 1` in `s`.

## Dry run (`aabxaab`)
- i=1: extend: s[0]=a = s[1]=a, s[1]=a ≠ s[2]=b → z=1; box [1,2)
- i=2: outside; s[0]=a ≠ b → 0
- i=3: x → 0
- i=4: extend a,a,b → 3; box [4,7)
- i=5: inside: min(7−5, z[1]=1) = 1; s[1]=a vs s[6]=b ✗ → 1
- i=6: inside: min(1, z[2]=0) = 0; s[0]=a vs b ✗ → 0
✅

- **Time:** O(n + m), **Space:** O(n + m)
