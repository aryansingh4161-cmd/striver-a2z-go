# Step 7: Recursion and Backtracking (Patterns)

> Status: ✅ theory and all 23 problems done
> Basics are in [Step 1.5](../step-01-basics/1.5-basic-recursion/).

## Pattern 1: Pick / not pick (subsequences)
For every element you make two choices: take it or skip it. That builds a binary recursion tree with `2ⁿ` leaves.

```go
func subsequences(a []int) [][]int {
    var res [][]int
    var path []int
    var f func(i int)
    f = func(i int) {
        if i == len(a) {
            res = append(res, append([]int{}, path...)) // save a COPY
            return
        }
        path = append(path, a[i]) // pick
        f(i + 1)
        path = path[:len(path)-1] // undo (backtrack)
        f(i + 1)                  // not pick
    }
    f(0)
    return res
}
```

Variations from this one template:
| Want | Change |
|---|---|
| print all subsequences with sum K | carry `sum`, check at the leaf |
| just **one** such subsequence | return `bool`; stop as soon as a call returns true |
| **count** them | return `int`: `return f(pick) + f(notPick)` |
| reuse an element (combination sum) | on "pick", stay at `i` instead of `i+1` |

## Pattern 2: Loop-based choices (avoid duplicate answers)
For "subsets II" / "combination sum II" (input has duplicates), instead of pick/not-pick, loop over the next choice and **skip equal values at the same level**:
```go
sort.Ints(a)
for j := i; j < len(a); j++ {
    if j > i && a[j] == a[j-1] { continue } // same choice at this level
    path = append(path, a[j])
    f(j + 1)
    path = path[:len(path)-1]
}
```

## Backtracking
Make a choice → recurse → **undo the choice**. Used for N-Queens, Sudoku, rat in a maze, word search, palindrome partitioning.
Always ask: what's the state, what are the choices, when is it invalid (prune early), when is it complete?

## Time complexity
Roughly **(number of choices)^(depth)** × work per leaf, e.g. O(2ⁿ × n) for subsequences.

## Go gotcha (again!)
Slices share memory. `res = append(res, path)` saves a reference that will keep changing. Always save `append([]int{}, path...)`.

## Problems

### 7.1 Build a strong base
| # | Problem | Key idea |
|---|---|---|
| 1 | [Recursive atoi](7.1-strong-base/01-recursive-atoi/) | parameterised recursion over digits |
| 2 | [Pow(x, n)](7.1-strong-base/02-pow/) | halve n each step |
| 3 | [Count good numbers](7.1-strong-base/03-count-good-numbers/) | 5^ceil(n/2) · 4^floor(n/2) with fast mod power |
| 4 | [Sort a stack](7.1-strong-base/04-sort-stack/) | pop, sort rest, insert in place |
| 5 | [Reverse a stack](7.1-strong-base/05-reverse-stack/) | pop, reverse rest, insert at bottom |

### 7.2 Subsequence pattern
| # | Problem | Key idea |
|---|---|---|
| 1 | [Binary strings without consecutive 1s](7.2-subsequences/01-binary-strings-no-consecutive-ones/) | add 1 only after 0 |
| 2 | [Generate parentheses](7.2-subsequences/02-generate-parentheses/) | open < n, close < open |
| 3 | [Power set / all subsequences](7.2-subsequences/03-power-set/) | pick / not pick, or bitmasks |
| 4 | [Count subsequences with sum K](7.2-subsequences/04-count-subsequences-sum-k/) | return f(pick) + f(skip) |
| 5 | [Any subsequence with sum K?](7.2-subsequences/05-subsequence-sum-k-exists/) | return true as soon as found |
| 6 | [Combination sum](7.2-subsequences/06-combination-sum/) | stay at i after picking (reuse) |
| 7 | [Combination sum II](7.2-subsequences/07-combination-sum-ii/) | sort + skip equal values per level |
| 8 | [Subset sums](7.2-subsequences/08-subset-sums/) | carry the sum, sort at the end |
| 9 | [Subsets II](7.2-subsequences/09-subsets-ii/) | every node is an answer + skip |
| 10 | [Combination sum III](7.2-subsequences/10-combination-sum-iii/) | loop d from start..9 |
| 11 | [Letter combinations of a phone number](7.2-subsequences/11-letter-combinations/) | one level per digit |

### 7.3 Hard
| # | Problem | Key idea |
|---|---|---|
| 1 | [Palindrome partitioning](7.3-hard/01-palindrome-partitioning/) | try every palindromic first piece |
| 2 | [Word search](7.3-hard/02-word-search/) | DFS, mark cell, unmark |
| 3 | [N-Queens](7.3-hard/03-n-queens/) | row by row, O(1) checks with col/diag arrays |
| 4 | [Rat in a maze](7.3-hard/04-rat-in-maze/) | DFS in D, L, R, U order |
| 5 | [M-colouring](7.3-hard/05-m-coloring/) | try each colour, check neighbours |
| 6 | [Sudoku solver](7.3-hard/06-sudoku-solver/) | row/col/box tables, backtrack |
| 7 | [Expression add operators](7.3-hard/07-expression-add-operators/) | track value and last term for * |
