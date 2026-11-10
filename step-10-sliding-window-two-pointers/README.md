# Step 10: Sliding Window and Two Pointers

> Status: ✅ theory and all 12 problems done

Used for **subarray / substring** problems: "longest", "shortest", "count", "max sum" of a contiguous part.

## Pattern 1: Fixed-size window (size k)
Add the new element on the right, remove the one falling off on the left.
```go
sum := 0
for i := 0; i < k; i++ { sum += a[i] }
best := sum
for r := k; r < len(a); r++ {
    sum += a[r] - a[r-k]
    best = max(best, sum)
}
```

## Pattern 2: Variable window: longest valid
Expand `r`; while the window is **invalid**, shrink from `l`. The window is valid after the inner loop → update the answer.
```go
l := 0
for r := 0; r < n; r++ {
    // add a[r] to the window state
    for windowInvalid() {
        // remove a[l] from the window state
        l++
    }
    best = max(best, r-l+1)
}
```
**Optimisation:** for "longest", you can shrink with `if` instead of `while` (the window never needs to get smaller than the best so far). Same answer, slightly faster.

## Pattern 3: Variable window: shortest valid
Expand `r`; **while** the window is valid, update the answer and shrink. (e.g. minimum window substring)

## Pattern 4: "Exactly K" = "at most K" − "at most K−1"
Counting subarrays with **exactly** K of something is hard directly, but **at most K** is a simple window: each `r` adds `r − l + 1` subarrays.
`exactly(K) = atMost(K) − atMost(K−1)`.
Used in: binary subarrays with sum, nice subarrays, subarrays with K different integers.

## Window state
Usually a counter: `map[byte]int`, `[26]int`, `[256]int`, or a "last seen index" array.

**Why O(n)?** `l` and `r` each move forward at most n times.

## Problems

### 10.1 Medium
| # | Problem | Key idea |
|---|---|---|
| 1 | [Longest substring without repeating characters](10.1-medium/01-longest-substring-no-repeat/) | jump l past last-seen copy |
| 2 | [Max consecutive ones III](10.1-medium/02-max-consecutive-ones-iii/) | at most k zeros in window |
| 3 | [Fruit into baskets](10.1-medium/03-fruit-into-baskets/) | at most 2 distinct |
| 4 | [Longest repeating character replacement](10.1-medium/04-longest-repeating-char-replacement/) | len − maxFreq ≤ k |
| 5 | [Binary subarrays with sum](10.1-medium/05-binary-subarrays-with-sum/) | atMost(g) − atMost(g−1) |
| 6 | [Count nice subarrays](10.1-medium/06-count-nice-subarrays/) | odd → 1, same as above |
| 7 | [Substrings containing all three characters](10.1-medium/07-substrings-all-three-chars/) | min(last) + 1 per end |
| 8 | [Maximum points from cards](10.1-medium/08-max-points-from-cards/) | slide the front/back split |

### 10.2 Hard
| # | Problem | Key idea |
|---|---|---|
| 1 | [Longest substring with at most K distinct](10.2-hard/01-longest-substring-k-distinct/) | shrink while > k distinct |
| 2 | [Subarrays with K different integers](10.2-hard/02-subarrays-k-different-integers/) | atMost(k) − atMost(k−1) |
| 3 | [Minimum window substring](10.2-hard/03-minimum-window-substring/) | shrink while valid |
| 4 | [Minimum window subsequence](10.2-hard/04-minimum-window-subsequence/) | forward match, backward tighten |
