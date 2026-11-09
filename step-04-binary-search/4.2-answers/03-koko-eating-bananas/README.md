# Koko Eating Bananas

**Problem (LeetCode 875):** `piles[i]` bananas in each pile. Each hour Koko picks one pile and eats up to `k` bananas from it (if the pile has fewer, she finishes it and waits for the rest of the hour). Find the **minimum speed `k`** to finish all piles within `h` hours.

`piles = [3 6 7 11]`, h = 8 → 4

## Checker
At speed `k`, a pile of `p` takes `ceil(p / k)` hours. Total hours = sum over piles.
`ceil(p/k)` in integers = `(p + k − 1) / k`.

## Search space
- `k = 1` is the slowest possible.
- `k = max(piles)` always works (one hour per pile).
- Faster speed → fewer hours → once a speed works, every bigger speed works. ✅ monotonic.

## Approach 1: Try every k from 1 upward
**O(max × n)**.

## Approach 2: Binary search on k (optimal)
Find the smallest `k` with `hours(k) <= h`.
- **Time:** O(n × log(max)), **Space:** O(1)

## Dry run `[3 6 7 11]`, h = 8
| k | hours | ≤ 8? |
|---|---|---|
| 6 (mid of 1..11) | 1+1+2+2 = 6 | ✅ → try smaller |
| 3 | 1+2+3+4 = 10 | ❌ |
| 4 | 1+2+2+3 = 8 | ✅ → answer **4** |
