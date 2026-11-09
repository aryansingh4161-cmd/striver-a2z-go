# Step 18: Strings (Advanced): Pattern Matching

> Status: theory ready, problems coming ⏳

**Problem:** find a pattern `p` (length m) in a text `s` (length n). Naive checking at every position is **O(n·m)**. These algorithms do it in **O(n + m)**.

## 1. Rabin-Karp (rolling hash)
Turn each window of length m into a number (hash). Sliding the window by one updates the hash in O(1): remove the left character, add the right one. Only compare characters when the hashes match.
- Average O(n + m). Use a large prime modulus to keep collisions rare.

## 2. Z-function
`z[i]` = length of the longest substring starting at `i` that matches a **prefix** of the string.
To find `p` in `s`, build the Z-array of `p + "$" + s`. Any position with `z[i] == m` is a match.
- O(n + m)

## 3. KMP (LPS array)
`lps[i]` = length of the longest **proper prefix** of `p[0..i]` that is also a **suffix** of it.
When a mismatch happens, the LPS array tells you how far you can jump back in the pattern **without moving back in the text**.
- O(n + m)

```go
func buildLPS(p string) []int {
    lps := make([]int, len(p))
    for i, length := 1, 0; i < len(p); {
        if p[i] == p[length] {
            length++
            lps[i] = length
            i++
        } else if length > 0 {
            length = lps[length-1] // fall back, don't move i
        } else {
            i++
        }
    }
    return lps
}
```

## Problems
| # | Problem | Status |
|---|---|---|
| 1 | Rabin-Karp: find all occurrences | ⏳ |
| 2 | Z-function | ⏳ |
| 3 | KMP / LPS array | ⏳ |
| 4 | Shortest palindrome (add characters in front) | ⏳ |
| 5 | Longest happy prefix | ⏳ |
| 6 | Count and say | ⏳ |
