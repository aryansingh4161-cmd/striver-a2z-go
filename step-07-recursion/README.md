# Step 7: Recursion and Backtracking (Patterns)

> Status: theory ready, problems coming ⏳
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
| # | Problem | Status |
|---|---|---|
| 1 | Recursive atoi | ⏳ |
| 2 | Pow(x, n) | ⏳ |
| 3 | Count good numbers | ⏳ |
| 4 | Sort a stack using recursion | ⏳ |
| 5 | Reverse a stack using recursion | ⏳ |

### 7.2 Subsequence pattern
| # | Problem | Status |
|---|---|---|
| 1 | Generate binary strings without consecutive 1s | ⏳ |
| 2 | Generate parentheses | ⏳ |
| 3 | Power set / all subsequences | ⏳ |
| 4 | Count subsequences with sum K | ⏳ |
| 5 | Check if a subsequence with sum K exists | ⏳ |
| 6 | Combination sum | ⏳ |
| 7 | Combination sum II | ⏳ |
| 8 | Subset sums | ⏳ |
| 9 | Subsets II | ⏳ |
| 10 | Combination sum III | ⏳ |
| 11 | Letter combinations of a phone number | ⏳ |

### 7.3 Hard
| # | Problem | Status |
|---|---|---|
| 1 | Palindrome partitioning | ⏳ |
| 2 | Word search | ⏳ |
| 3 | N-Queens | ⏳ |
| 4 | Rat in a maze | ⏳ |
| 5 | M-colouring problem | ⏳ |
| 6 | Sudoku solver | ⏳ |
| 7 | Expression add operators | ⏳ |
