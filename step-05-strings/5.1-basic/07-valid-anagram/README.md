# Valid Anagram

**Problem (LeetCode 242):** Is `t` a rearrangement of `s`? (lowercase letters)

`"anagram"`, `"nagaram"` → true

## Approach 1: Sort both
Sorted strings must be equal. **O(n log n)**.

## Approach 2: Frequency array (optimal)
`count[26]`: `+1` for each char of `s`, `−1` for each of `t`. All zeros → anagram.
- **Time:** O(n), **Space:** O(1) (26 counters)

For Unicode text, use a `map[rune]int` instead of the array.
