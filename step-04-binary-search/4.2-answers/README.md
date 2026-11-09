# 4.2 Binary Search on Answers

The array isn't searched directly. Instead you search over the **possible answers**.

## When does it apply?
The question looks like "find the **minimum** X such that something is possible" (or the **maximum**), and:
- if X works, every **bigger** X also works (for "minimum" questions), or
- if X works, every **smaller** X also works (for "maximum" questions).

So the answers look like: ❌ ❌ ❌ ✅ ✅ ✅ ✅. Binary search finds the first ✅.

## The recipe
1. **Range:** smallest possible answer `lo`, largest possible answer `hi`.
2. **Checker:** write `possible(x) bool` (usually a simple O(n) loop).
3. **Search:**
```go
ans := hi
for lo <= hi {
    mid := lo + (hi-lo)/2
    if possible(mid) {
        ans = mid      // works; try smaller
        hi = mid - 1
    } else {
        lo = mid + 1
    }
}
return ans
```
For "maximum X that works", flip it: on success save and go **right**.

**Time:** O(log(range) × cost of `possible`).

## Problems
| # | Problem | Search space | Checker |
|---|---|---|---|
| 1 | [Square root](01-square-root/) | 1..n | `mid*mid <= n` |
| 2 | [Nth root](02-nth-root/) | 1..m | `mid^n` vs m |
| 3 | [Koko eating bananas](03-koko-eating-bananas/) | speed 1..max pile | hours needed ≤ h |
| 4 | [Min days for M bouquets](04-min-days-bouquets/) | day min..max | bouquets ≥ m |
| 5 | [Smallest divisor](05-smallest-divisor/) | 1..max | sum of ceilings ≤ threshold |
| 6 | [Ship packages in D days](06-ship-packages/) | max weight..total | days needed ≤ D |
| 7 | [Kth missing positive](07-kth-missing-positive/) | index | missing count before index |
| 8 | [Aggressive cows](08-aggressive-cows/) | distance 1..range | can place all cows |
| 9 | [Book allocation](09-book-allocation/) | max..total pages | students needed ≤ m |
| 10 | [Split array largest sum](10-split-array-largest-sum/) | same as above | parts needed ≤ k |
| 11 | [Painter's partition](11-painters-partition/) | same as above | painters needed ≤ k |
| 12 | [Minimise max gas station distance](12-minimise-max-distance-gas-station/) | real numbers | stations needed ≤ k |
| 13 | [Median of two sorted arrays](13-median-two-sorted-arrays/) | how many from array 1 | valid partition |
| 14 | [Kth element of two sorted arrays](14-kth-element-two-sorted-arrays/) | how many from array 1 | valid partition |
