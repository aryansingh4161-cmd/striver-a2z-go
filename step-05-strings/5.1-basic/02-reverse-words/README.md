# Reverse Words in a String

**Problem (LeetCode 151):** Reverse the **order of words**. Remove leading/trailing spaces and collapse multiple spaces into one.

`"  the sky   is blue "` → `"blue is sky the"`

## Approach 1: Split + reverse (easiest in Go)
`strings.Fields(s)` splits on any run of spaces and drops empty pieces. Reverse the slice, then `strings.Join(words, " ")`.
- **Time:** O(n), **Space:** O(n)

## Approach 2: Two pointers from the end (no library split)
Walk from the right. Skip spaces; mark where a word ends; walk left to its start; append that word. This is what you'd write in an interview that bans library helpers.
- **Time:** O(n), **Space:** O(n) for the answer

(In C++ the "O(1) extra space" version reverses the whole string, then reverses each word. Go strings are immutable, so you'd need a `[]byte` copy anyway.)
