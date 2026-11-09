# Remove Outermost Parentheses

**Problem (LeetCode 1021):** A valid parentheses string is made of "primitive" parts (valid, can't be split further). Remove the outermost pair of **each** primitive part.

`"(()())(())"` → `"()()()"`
Primitives: `(()())` → `()()`, and `(())` → `()`.

## Approach: Depth counter (optimal)
Keep `depth` = how many `(` are currently open.
- On `(`: if `depth > 0`, it's not outermost → keep it. Then `depth++`.
- On `)`: `depth--` first. If `depth > 0` after that, keep it.

The outermost `(` is the one seen at depth 0; the outermost `)` is the one that brings depth back to 0.
- **Time:** O(n), **Space:** O(n) for the result

## Dry run `"(()())"`
| ch | depth before | keep? | depth after |
|---|---|---|---|
| ( | 0 | no (outer) | 1 |
| ( | 1 | yes | 2 |
| ) | 2 | → 1, yes | 1 |
| ( | 1 | yes | 2 |
| ) | 2 | → 1, yes | 1 |
| ) | 1 | → 0, no (outer) | 0 |
Result `()()`.
