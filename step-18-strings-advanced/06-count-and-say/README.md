# Count and Say

**Problem:** The sequence starts with `"1"`. Each next term **reads the previous term aloud**: group equal consecutive digits and say "count, digit".

```
1           → "one 1"            → 11
11          → "two 1s"           → 21
21          → "one 2, one 1"     → 1211
1211        → "one 1, one 2, two 1s" → 111221
111221      → 312211
```
Return the `n`-th term (1-indexed).

## Approach: simulate with run-length encoding
For each step, scan the current string, counting the length of each run of equal digits, and write `count` then `digit`.

```
cur = "1"
repeat n−1 times:
    next = ""
    i = 0
    while i < len(cur):
        j = i
        while j < len(cur) and cur[j] == cur[i]: j++
        next += str(j − i) + cur[i]
        i = j
    cur = next
```
Use a `strings.Builder` in Go so building each term is linear, not quadratic.

## Dry run (term 4 → 5)
`1211`: runs `1` (×1), `2` (×1), `11` (×2) → "11" + "12" + "21" = **111221** ✅

## Complexity
The terms grow by about 30% each step (Conway's constant ≈ 1.3036), so the work is the total length of all terms: roughly O(1.3ⁿ). For the usual limit `n ≤ 30` that's tiny.

## Fun fact
Starting from "1", no digit other than 1, 2, 3 ever appears.
