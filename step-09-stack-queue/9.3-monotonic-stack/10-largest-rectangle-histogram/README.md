# Largest Rectangle in a Histogram

**Problem (LeetCode 84):** Bars of width 1 and heights `h`. Find the largest rectangle area.

`[2 1 5 6 2 3]` → 10 (bars 5 and 6, height 5, width 2)

## Key insight
The best rectangle uses some bar `i` as its **height** (the shortest bar inside it). It stretches left and right until it hits a **shorter** bar.
Width = `NSE index − PSE index − 1`.

## Approach 1: For each bar, expand both ways
**O(n²)**.

## Approach 2: PSE and NSE arrays with stacks
Two passes to find previous/next smaller **indexes**, then one pass for areas. **Time O(n), Space O(n)**.

## Approach 3: One-pass stack (optimal)
Keep indexes with **increasing** heights. When a shorter bar `i` arrives, every taller bar on top has just found its **next smaller** (`i`); its **previous smaller** is the element below it in the stack. Pop and compute its area right then.
At the end, process remaining bars with "next smaller" = `n`.
```
for i in 0..n:   (h = 0 at i == n flushes the stack)
    while stack not empty and h[top] >= cur:
        height = h[pop]
        left = new top or -1
        area = height * (i - left - 1)
    push i
```
- **Time:** O(n), **Space:** O(n)

## Dry run `[2 1 5 6 2 3]`
At i=4 (height 2): pop 6 → width 4−2−1 = 1, area 6; pop 5 → width 4−1−1 = 2, area **10**.
