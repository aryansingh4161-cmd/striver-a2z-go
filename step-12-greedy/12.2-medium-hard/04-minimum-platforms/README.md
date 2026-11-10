# Minimum Number of Platforms

**Problem:** Train arrival and departure times. A train needs a platform from arrival to departure (inclusive: if one arrives at the same time another departs, they need **different** platforms). Minimum platforms so no train waits?

arr `[900 940 950 1100 1500 1800]`, dep `[910 1200 1120 1130 1900 2000]` → 3

## Approach 1: Count overlaps for each train
For each train, count how many others overlap it. **O(n²)**.

## Approach 2: Sort arrivals and departures separately, sweep (optimal)
The answer is the **maximum number of trains present at the same time**.
Sort both lists. Walk with two pointers like merging:
- next event is an **arrival** (`arr[i] <= dep[j]`) → `count++`, `i++`
- otherwise a **departure** → `count--`, `j++`
Track the max count.

`<=` means an arrival at the same minute as a departure is processed **first**: both trains need a platform at that moment.
- **Time:** O(n log n), **Space:** O(1) besides sorting

## Dry run (times sorted)
arrivals 900 940 950 1100 1500 1800, departures 910 1120 1130 1200 1900 2000
900 A(1) · 910 D(0) · 940 A(1) · 950 A(2) · 1100 A(**3**) · 1120 D(2) · 1130 D(1) · 1200 D(0) · …
