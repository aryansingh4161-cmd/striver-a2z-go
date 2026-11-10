# Generate Binary Strings Without Consecutive 1s

**Problem:** Return all binary strings of length `n` with no two adjacent `1`s, in sorted order.

n = 3 → `000 001 010 100 101`

## Approach: Recursion with a rule
Build the string one character at a time:
- you can **always** add `0`,
- you can add `1` **only if the last character isn't `1`**.

```
f(cur):
    if len(cur) == n: save cur; return
    f(cur + "0")
    if cur is empty or last char is '0': f(cur + "1")
```
Trying `0` before `1` gives sorted output.

- **Time:** O(number of valid strings × n). That count is a Fibonacci number (≈ 1.618ⁿ), far fewer than 2ⁿ.
- **Space:** O(n) recursion depth

This is the simplest **"choices with a constraint"** recursion: the same shape as every backtracking problem in this step.
