# Matrix Chain Multiplication (MCM)

**Problem:** `arr` has length `n`; matrix `i` (for `i = 1..n−1`) has size `arr[i−1] × arr[i]`. Multiplying a `p × q` by a `q × r` matrix costs `p · q · r` scalar multiplications. Find the **minimum** total cost to multiply the whole chain (you choose where the brackets go).

```
arr = [10 20 30 40 50] → 4 matrices: A(10×20) B(20×30) C(30×40) D(40×50)
best: ((AB)C)D = 6000 + 12000 + 20000 = 38000
```

## The partition DP pattern (this whole section)
1. The problem is about a **range** `[i, j]`.
2. Try **every split point** `k` in the range: solve the left part and the right part separately, plus the cost of combining them.
3. Take the best.

```
f(i, j) = 0                                       if i == j  (one matrix: nothing to multiply)
f(i, j) = min over k = i..j−1 of
          f(i, k) + f(k+1, j) + arr[i−1] · arr[k] · arr[j]
answer  = f(1, n−1)
```
`arr[i−1] · arr[k] · arr[j]` is the cost of multiplying the left result (`arr[i−1] × arr[k]`) by the right result (`arr[k] × arr[j]`).

## Stages
- Recursion: exponential (Catalan number of bracketings)
- Memo: n² states × n splits → **O(n³)**, Space O(n²)
- Tabulation: `i` goes from **n−1 down to 1**, `j` from **i+1 up to n−1** (so the smaller ranges inside are ready).

## Dry run (`[10 20 30 40]`, matrices 1..3)
- f(1,2) = 10·20·30 = 6000; f(2,3) = 20·30·40 = 24000
- f(1,3) = min( k=1: 0 + 24000 + 10·20·40 = 32000,
                k=2: 6000 + 0 + 10·30·40 = 18000 ) = **18000** ✅
