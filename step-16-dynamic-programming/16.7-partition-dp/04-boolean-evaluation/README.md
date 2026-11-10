# Evaluate Boolean Expression to True

**Problem:** A string of `T` / `F` separated by operators `&`, `|`, `^`, e.g. `"T|T&F^T"`. Count the number of ways to put brackets so the whole expression evaluates to **true** (mod `1e9 + 7`).

```
"T|T&F^T" → 4
"T^F|F"   → 2   ( (T^F)|F and T^(F|F) )
```

## Partition at every operator
Symbols are at even indexes, operators at odd indexes. For a range `[i, j]` (both symbols), split at every operator `k` (odd index between them): left is `[i, k−1]`, right is `[k+1, j]`.

To combine, we need **both** how many ways each side is true **and** false:
| op | ways to be true | ways to be false |
|---|---|---|
| `&` | lT·rT | lT·rF + lF·rT + lF·rF |
| `\|` | lT·rT + lT·rF + lF·rT | lF·rF |
| `^` | lT·rF + lF·rT | lT·rT + lF·rF |

So the state is `f(i, j, wantTrue)`.
```
f(i, i, true)  = 1 if s[i] == 'T'
f(i, i, false) = 1 if s[i] == 'F'
f(i, j, want)  = Σ over operators k: combine(f(i,k−1,·), f(k+1,j,·), s[k], want)
```

## Stages
- Recursion: exponential
- Memo / table `dp[i][j][2]`: **O(n³)**, Space O(n²)

## Dry run (`T^F|F`)
- split at `^`: left `T` (T=1, F=0), right `F|F` (T=0, F=1) → `^` true = 1·1 + 0·0 = **1**
- split at `|`: left `T^F` (T=1, F=0), right `F` (T=0, F=1) → `|` true = 0 + 1·1 + 0 = **1**
- total **2** ✅
