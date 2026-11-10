# Print the LCS

**Problem:** Not just the length: return one actual longest common subsequence string.

## Idea: walk the LCS table backwards
Build the [LCS table](../01-lcs/) `dp`. Start at the bottom-right `(n, m)` and retrace the decisions:
```
i, j = n, m
while i > 0 and j > 0:
    if a[i−1] == b[j−1]:       // this char is in the LCS
        add a[i−1]; i--; j--
    else if dp[i−1][j] >= dp[i][j−1]:
        i--                    // the answer came from above
    else:
        j--                    // the answer came from the left
reverse what you collected
```
We collect characters from the end, so reverse at the end (or fill a byte slice of length `dp[n][m]` from the back).

## Dry run (`abcde` / `bdgek`)
LCS table (rows = prefixes of `abcde`, columns = prefixes of `bdgek`):
```
      ""  b  d  g  e  k
""     0  0  0  0  0  0
a      0  0  0  0  0  0
b      0  1  1  1  1  1
c      0  1  1  1  1  1
d      0  1  2  2  2  2
e      0  1  2  2  3  3
```
Walk back from (5, 5):
| at | compare | move |
|---|---|---|
| (5,5) | e vs k differ; up 2 < left 3 | left |
| (5,4) | e = e | take **e**, diagonal |
| (4,3) | d vs g differ; up 1 < left 2 | left |
| (4,2) | d = d | take **d**, diagonal |
| (3,1) | c vs b differ; up 1 ≥ left 0 | up |
| (2,1) | b = b | take **b** → done |
Collected backwards: e, d, b → **"bde"** ✅

There can be several LCS strings of the same length; this returns one of them (which one depends on the tie-break `>=`).

- **Time:** O(n · m) to build + O(n + m) to walk, **Space:** O(n · m)
