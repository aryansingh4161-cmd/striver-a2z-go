# Step 5: Strings (Basic and Medium)

> Status: theory ready, problems coming ⏳

## Strings in Go: what you must know
- A `string` is a **read-only** sequence of bytes. `s[i]` is a `byte`.
- To change characters, convert: `b := []byte(s)`, edit `b`, then `string(b)`.
- `len(s)` = number of **bytes**. For English letters, bytes = characters. For other text use `[]rune(s)`.
- Building a string in a loop with `+=` copies everything each time (O(n²) total). Use `strings.Builder`.
- Comparing strings with `==` and `<` works (dictionary order).

```go
var sb strings.Builder
for i := len(s) - 1; i >= 0; i-- {
    sb.WriteByte(s[i])
}
reversed := sb.String()
```

Handy functions: `strings.Fields` (split on spaces), `strings.Join`, `strings.HasPrefix`, `strings.Repeat`, `strconv.Atoi`, `strconv.Itoa`.

## Common patterns
| Pattern | Used for |
|---|---|
| Frequency array `[26]int` | anagrams, sort by frequency, beauty of substrings |
| Two pointers from both ends | palindromes, reversing words |
| Map char → char | isomorphic strings |
| Counter for depth | parentheses problems |
| Expand around centre | longest palindromic substring |
| Sliding window | substrings with K distinct characters (also Step 10) |

## Problems

### 5.1 Basic
| # | Problem | Status |
|---|---|---|
| 1 | Remove outermost parentheses | ⏳ |
| 2 | Reverse words in a string | ⏳ |
| 3 | Largest odd number in a string | ⏳ |
| 4 | Longest common prefix | ⏳ |
| 5 | Isomorphic strings | ⏳ |
| 6 | Check whether one string is a rotation of another | ⏳ |
| 7 | Valid anagram | ⏳ |

### 5.2 Medium
| # | Problem | Status |
|---|---|---|
| 1 | Sort characters by frequency | ⏳ |
| 2 | Maximum nesting depth of parentheses | ⏳ |
| 3 | Roman to integer / integer to Roman | ⏳ |
| 4 | Implement atoi | ⏳ |
| 5 | Count substrings with exactly K distinct characters | ⏳ |
| 6 | Longest palindromic substring | ⏳ |
| 7 | Sum of beauty of all substrings | ⏳ |
| 8 | Reverse every word in a string | ⏳ |
