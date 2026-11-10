# Count Subsets with Sum K

**Problem:** How many subsets of `arr` have sum exactly `k`? (Elements are non-negative; answer mod `1e9 + 7`.)

```
arr = [1 2 2 3], k = 3 → 3 ({1, 2}, {1, 2'}, {3})
```

## From "is it possible?" to "how many?"
Same take / not take as [Subset Sum](../01-subset-sum/), but **add** the two counts instead of OR-ing them:
```
f(i, t) = f(i−1, t)                         // not take
        + (arr[i] <= t ? f(i−1, t − arr[i]) : 0)   // take
```

## The base case with zeros (common bug!)
If you use `f(i, 0) = 1` ("sum reached, stop"), you **miss** subsets that also include zeros. With `arr = [0 0 1]`, `k = 1`, the right answer is **4** (`{1}`, `{0,1}`, `{0',1}`, `{0,0',1}`).

The safe way: go all the way down to index 0 and handle it carefully:
```
f(0, t):
    if t == 0 and arr[0] == 0: return 2   // take it or not, both give 0
    if t == 0 or t == arr[0]:  return 1
    return 0
```
Then `f(i, 0)` is **not** a base case anymore; it's computed like everything else.

## Stages
- Recursion: O(2ⁿ)
- Memo / table: O(n · k)
- One array (loop `t` downward): **O(k)** space

## Dry run (`[1 2 2 3]`, k = 3), one array `ways[t]`
| after | t=0 | 1 | 2 | 3 |
|---|---|---|---|---|
| start | 1 | 0 | 0 | 0 |
| 1 | 1 | 1 | 0 | 0 |
| 2 | 1 | 1 | 1 | 1 |
| 2 | 1 | 1 | 2 | 2 |
| 3 | 1 | 1 | 2 | **3** |
(Starting the 1D array at `ways[0] = 1` and looping over **all** elements handles zeros correctly: a zero doubles `ways[0]`.)
