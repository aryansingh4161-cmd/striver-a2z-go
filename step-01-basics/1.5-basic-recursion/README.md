# 1.5 Basic Recursion

**Recursion** = a function that calls itself to solve a smaller version of the same problem.

## The 3 parts of every recursive function

```go
func f(n int) {
    if n == 0 {      // 1. BASE CASE: when to stop
        return
    }
    // 2. WORK for this level
    f(n - 1)         // 3. RECURSIVE CALL on a SMALLER input
}
```

Forget the base case (or never get closer to it) → the function calls itself forever → **stack overflow** (Go prints `goroutine stack exceeds limit`).

## The call stack

Each call waits on a stack until the call it made returns. `f(3)` → `f(2)` → `f(1)` → `f(0)` returns → `f(1)` finishes → … → `f(3)` finishes.

- Depth `n` → **O(n) stack space**, even if you don't create any slices.

## Recursion tree

Draw each call as a node and its calls as children. It shows you:
- the **order** things happen in,
- the **number of calls** (that gives the time complexity).

Fibonacci `f(4)`:
```
            f(4)
          /      \
       f(3)      f(2)
      /   \      /   \
   f(2)  f(1)  f(1) f(0)
   /  \
 f(1) f(0)
```
Two calls per level → roughly **O(2ⁿ)** calls.

## Head vs tail work: printing 1..N vs N..1

| Code | Prints (n = 3) |
|---|---|
| print **before** the call: `print(n); f(n-1)` | 3 2 1 |
| print **after** the call: `f(n-1); print(n)` | 1 2 3 |

Work done after the call happens while the stack **unwinds** (backtracking). This idea is used everywhere later.

## Two styles

| Style | How | Example: sum of 1..n |
|---|---|---|
| **Parameterised** | carry the answer in a parameter | `f(i, sum)`: at the end, `sum` holds the answer |
| **Functional** | the function **returns** the answer | `return n + f(n-1)` |

Functional style is more common in later topics (trees, DP).

## How to think recursively
1. Trust that `f(n-1)` already works (the "leap of faith").
2. Ask: how do I use `f(n-1)` to build `f(n)`?
3. Write the smallest case directly (base case).

## Problems
| # | Problem | Idea |
|---|---|---|
| 1 | [Print name N times](01-print-name-n-times/) | count down to 0 |
| 2 | [Print 1 to N](02-print-1-to-n/) | work after the call (backtracking) |
| 3 | [Print N to 1](03-print-n-to-1/) | work before the call |
| 4 | [Sum of first N](04-sum-first-n/) | parameterised vs functional |
| 5 | [Factorial](05-factorial/) | `n * f(n-1)` |
| 6 | [Reverse array](06-reverse-array/) | swap ends, recurse inward |
| 7 | [Palindrome string](07-palindrome-string/) | compare ends, recurse inward |
| 8 | [Fibonacci](08-fibonacci/) | multiple recursive calls |

Note: the "print" problems **return a slice** instead of printing, so they can be tested.
