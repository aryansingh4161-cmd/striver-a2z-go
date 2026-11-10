# Surrounded Regions

**Problem (LeetCode 130):** Board of `'X'` and `'O'`. Capture every region of `'O'` that is **completely surrounded** by `'X'` (flip it to `'X'`). An `'O'` region touching the **border** can't be captured.

```
X X X X        X X X X
X O O X   →    X X X X
X X O X        X X X X
X O X X        X O X X     (bottom O touches the border)
```

## Approach: Think in reverse: save the border-connected O's
1. DFS from every `'O'` on the **border**, marking every `'O'` reachable from it as **safe**.
2. Every `'O'` not marked safe is surrounded → flip to `'X'`.
- **Time:** O(n × m), **Space:** O(n × m)

Finding "what's surrounded" directly is awkward; finding "what escapes" from the border is easy.
