# Recursive atoi

**Problem:** Same rules as [atoi in Step 5](../../../step-05-strings/5.2-medium/04-atoi/) (skip spaces, optional sign, read digits, clamp to 32-bit), but read the digits **with recursion** instead of a loop.

`"   -42abc"` → −42

## Approach
1. Skip spaces and read the sign with a normal loop (that part isn't interesting).
2. `digits(i, acc)`:
   - **base case:** `i` is past the end or `s[i]` isn't a digit → return `acc`
   - **work:** `acc = acc*10 + digit`
   - **clamp early:** if `acc` already passed 2³¹, return it (the caller clamps)
   - **recurse:** `digits(i+1, acc)`

This is **parameterised recursion** (see [Step 1.5](../../../step-01-basics/1.5-basic-recursion/)): the answer is carried in `acc`.
- **Time:** O(n), **Space:** O(n) recursion stack
