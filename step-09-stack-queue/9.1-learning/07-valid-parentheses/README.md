# Valid Parentheses

**Problem (LeetCode 20):** A string of `()[]{}`. Valid if every opener is closed by the **same type** in the **correct order**.

`"()[]{}"` → true, `"(]"` → false, `"([)]"` → false, `"{[]}"` → true

## Approach: Stack of openers (optimal)
- Opener → push it.
- Closer → the stack must not be empty, and its top must be the **matching** opener; pop it.
- At the end the stack must be **empty** (otherwise something wasn't closed).

- **Time:** O(n), **Space:** O(n)

## Dry run `"{[]}"`
| ch | stack after |
|---|---|
| { | { |
| [ | { [ |
| ] | { (matched [) |
| } | (empty, matched {) → **true** |

Only one type of bracket? A simple counter is enough (see [nesting depth](../../../step-05-strings/5.2-medium/02-max-nesting-depth/)). Mixed types need the stack, to remember the **order**.
