# Implement a Stack Using an Array

**Stack = LIFO.** Operations: `Push(x)`, `Pop()`, `Top()`, `Size()`, `IsEmpty()`.

## Fixed-size array version (what Striver shows)
Keep an array of capacity `cap` and an index `top` (−1 = empty).
- **Push:** if `top == cap−1` → overflow; else `top++`, `arr[top] = x`
- **Pop:** if `top == −1` → underflow; else return `arr[top]`, `top--`
- **Top:** `arr[top]`
- All **O(1)**

## Go slice version (what you'd use)
A slice grows automatically, so there's no overflow:
`append` to push, `s[len(s)−1]` to peek, `s = s[:len(s)−1]` to pop. All O(1) (append is O(1) amortised).

Both are implemented in `solution.go`. Errors are returned as `(value, ok)` instead of crashing.
