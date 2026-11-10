# Maximum Points You Can Obtain from Cards

**Problem (LeetCode 1423):** Take exactly `k` cards, each from the **front or the back** of the row. Maximise the total.

`[1 2 3 4 5 6 1]`, k = 3 → 12 (take the last three: 1 + 6 + 5)

## Key insight
Any choice = take `i` cards from the front and `k − i` from the back, for some `i` in `0..k`.

## Approach: Slide the split point (optimal)
1. Start with all `k` from the front: `sum = a[0] + … + a[k−1]`.
2. Move one card at a time from the front to the back: remove `a[i]` (from the end of the front part), add the next card from the back. Track the max.
- **Time:** O(k), **Space:** O(1)

Equivalent view: you leave a **contiguous window** of `n − k` cards in the middle; minimise that window's sum.
