# Step 5: Strings (Basic and Medium)

> Status: ✅ theory and all 15 problems done

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Remove outermost parentheses](5.1-basic/01-remove-outermost-parentheses/) | depth counter |
| 2 | [Reverse words in a string](5.1-basic/02-reverse-words/) | scan from the end / Fields |
| 3 | [Largest odd number in a string](5.1-basic/03-largest-odd-number/) | rightmost odd digit |
| 4 | [Longest common prefix](5.1-basic/04-longest-common-prefix/) | column scan / sort first & last |
| 5 | [Isomorphic strings](5.1-basic/05-isomorphic-strings/) | last-seen positions |
| 6 | [Rotate string](5.1-basic/06-rotate-string/) | goal inside s+s |
| 7 | [Valid anagram](5.1-basic/07-valid-anagram/) | 26 counters |

### 5.2 Medium
| # | Problem | Key idea |
|---|---|---|
| 1 | [Sort characters by frequency](5.2-medium/01-sort-characters-by-frequency/) | count + sort |
| 2 | [Maximum nesting depth](5.2-medium/02-max-nesting-depth/) | running depth |
| 3 | [Roman ↔ integer](5.2-medium/03-roman-integer/) | subtract if smaller than next; greedy table |
| 4 | [String to integer (atoi)](5.2-medium/04-atoi/) | step-by-step + clamp |
| 5 | [Count substrings with exactly K distinct](5.2-medium/05-count-substrings-k-distinct/) | atMost(k) − atMost(k−1) |
| 6 | [Longest palindromic substring](5.2-medium/06-longest-palindromic-substring/) | expand around centre |
| 7 | [Sum of beauty of all substrings](5.2-medium/07-sum-of-beauty/) | running frequency |
| 8 | [Reverse each word](5.2-medium/08-reverse-each-word/) | two pointers per word |
