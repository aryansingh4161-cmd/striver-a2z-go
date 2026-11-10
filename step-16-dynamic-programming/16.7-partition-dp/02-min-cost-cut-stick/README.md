# Minimum Cost to Cut a Stick

**Problem:** A stick of length `n` with marks at positions `cuts`. Cutting a piece costs **its current length**. You can do the cuts in any order. Minimum total cost?

```
n = 7, cuts = [1 3 4 5] → 16
order 3, 5, 1, 4: 7 + 4 + 3 + 2 = 16
```

## Turning it into MCM
1. **Sort** `cuts` so that every piece between two cuts is a contiguous range of cut indexes.
2. Add the stick ends: `c = [0, cuts..., n]`.
3. `f(i, j)` = min cost to make all cuts `c[i..j]` (indexes `1..m`), inside the piece from `c[i−1]` to `c[j+1]`.
   - Choose which cut `k` in `i..j` to do **first**: it costs the whole piece `c[j+1] − c[i−1]`, then the two sides are independent.
```
f(i, j) = 0                                         if i > j
f(i, j) = min over k = i..j of
          (c[j+1] − c[i−1]) + f(i, k−1) + f(k+1, j)
answer = f(1, m)
```

## Why sorting is required
Without it, "cuts `i..j`" wouldn't describe a single piece of stick, and the two halves after a cut wouldn't be independent ranges.

## Stages
- Recursion: exponential
- Memo / table: m² states × m choices → **O(m³)**, Space O(m²)

## Dry run (`n = 7`, `c = [0 1 3 4 5 7]`)
Doing cut at 3 first on the whole stick costs 7; then left piece [0, 3] has cut 1 (cost 3), right piece [3, 7] has cuts 4, 5 (cut 5 first: 4, then 4: 2) → 7 + 3 + 4 + 2 = **16** ✅
