# Sum of First N Numbers

**Problem:** Return 1 + 2 + … + N.

| N | Output |
|---|---|
| 3 | 6 |
| 5 | 15 |

## Approach 1: Parameterised recursion
Carry the running total in a parameter.
```
f(i, sum):
    if i < 1: answer = sum; return
    f(i - 1, sum + i)
```

## Approach 2: Functional recursion
Return the answer.
```
f(n):
    if n == 0: return 0
    return n + f(n - 1)
```
Trace: `f(3) = 3 + f(2) = 3 + 2 + f(1) = 3 + 2 + 1 + f(0) = 6`.

Both: **Time O(n), Space O(n)** stack.

## Approach 3: Formula
`N × (N + 1) / 2`: **O(1)**. Use this in real code. The recursive versions are for practising recursion.
