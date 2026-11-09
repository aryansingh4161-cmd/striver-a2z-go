# Largest Odd Number in a String

**Problem (LeetCode 1903):** `num` is a string of digits. Return the largest-valued odd number that is a **substring** (contiguous), or `""`.

`"52"` → `"5"`, `"35427"` → `"35427"`, `"4206"` → `""`

## Key insight
A number is odd iff its **last digit** is odd. The longest substring starting at index 0 is always the largest. So: find the **rightmost odd digit** and return everything up to it.
- **Time:** O(n), **Space:** O(1) (returns a substring of the input)

(If the string can have leading zeros like `"0035"`, skip them first. LeetCode guarantees no leading zeros; the solution handles it anyway.)
