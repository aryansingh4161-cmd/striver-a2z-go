# Rotate String

**Problem (LeetCode 796):** Can `s` become `goal` after some number of left shifts (move the first character to the end)?

`"abcde"`, `"cdeab"` → true

## Approach 1: Try every rotation
For each k, compare `s[k:] + s[:k]` with `goal`. **O(n²)**.

## Approach 2: Doubling trick (optimal idea)
Every rotation of `s` appears inside `s + s`:
`"abcdeabcde"` contains `"cdeab"`, `"deabc"`, …
So: same length **and** `goal` is a substring of `s + s`.
- **Time:** O(n) with a linear-time substring search (KMP, Step 18). Go's `strings.Contains` is fast in practice.
- **Space:** O(n)
