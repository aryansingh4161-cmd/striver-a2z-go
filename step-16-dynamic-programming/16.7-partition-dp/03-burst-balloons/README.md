# Burst Balloons

**Problem:** Balloons with numbers `nums`. Bursting balloon `i` gives `nums[left] · nums[i] · nums[right]` coins, where `left` and `right` are its **current** neighbours (out of bounds counts as 1). Burst all balloons; maximise coins.

```
[3 1 5 8] → 167
burst 1: 3·1·5 = 15 → [3 5 8]
burst 5: 3·5·8 = 120 → [3 8]
burst 3: 1·3·8 = 24 → [8]
burst 8: 1·8·1 = 8 → total 167
```

## Why "first to burst" doesn't work
If you burst `k` **first**, its neighbours change, and the left and right parts are **not independent** (balloons on the left will later be neighbours of balloons on the right).

## The trick: choose the **last** balloon to burst in the range
If `k` is the **last** balloon burst in range `[i, j]`, then while everything else in the range is being burst, `k` is still standing and acts as a **wall**. So `[i, k−1]` and `[k+1, j]` are independent. When `k` finally bursts, its neighbours are whatever is just outside the range: `a[i−1]` and `a[j+1]`.

Pad with 1s: `a = [1, nums..., 1]`.
```
f(i, j) = 0                                           if i > j
f(i, j) = max over k = i..j of
          a[i−1] · a[k] · a[j+1] + f(i, k−1) + f(k+1, j)
answer = f(1, n)
```

## Stages
- Recursion: exponential
- Memo / table: **O(n³)** time, O(n²) space (fill `i` from n down to 1, `j` from i up to n)

## Dry run (`[3 1 5 8]`, a = `[1 3 1 5 8 1]`)
The best plan has **8** last (`1 · 8 · 1 = 8`), with `[3 1 5]` done before it with walls 1 and 8. Inside that, **3** is last (`1 · 3 · 8 = 24`), and before it `[1 5]` with walls 3 and 8: 5 last (`3 · 5 · 8 = 120`), 1 first (`3 · 1 · 5 = 15`). Total `8 + 24 + 120 + 15 = 167` ✅
