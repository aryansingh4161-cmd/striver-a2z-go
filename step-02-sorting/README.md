# Step 2: Sorting Techniques

Sorting = arranging elements in order (we use ascending). Learning the algorithms teaches loop control, recursion and divide & conquer, which are used everywhere later.

## Summary table

| Algorithm | Idea in one line | Best | Average | Worst | Extra space | Stable? |
|---|---|---|---|---|---|---|
| [Selection](01-selection-sort/) | pick the minimum, put it in front | O(n²) | O(n²) | O(n²) | O(1) | No |
| [Bubble](02-bubble-sort/) | swap neighbours, biggest bubbles to the end | O(n)* | O(n²) | O(n²) | O(1) | Yes |
| [Insertion](03-insertion-sort/) | insert each element into the sorted left part | O(n) | O(n²) | O(n²) | O(1) | Yes |
| [Merge](04-merge-sort/) | split in half, sort halves, merge | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| [Recursive bubble](05-recursive-bubble-sort/) | one bubble pass, recurse on n-1 | O(n)* | O(n²) | O(n²) | O(n) stack | Yes |
| [Recursive insertion](06-recursive-insertion-sort/) | sort n-1, insert the last | O(n) | O(n²) | O(n²) | O(n) stack | Yes |
| [Quick](07-quick-sort/) | place a pivot, sort both sides | O(n log n) | O(n log n) | O(n²) | O(log n) stack | No |

\* with the "no swaps → stop early" check.

**Stable** = equal elements keep their original order. Matters when sorting records by one field (e.g. students by marks).

## Which one to remember for interviews?
- **Merge sort**: guaranteed O(n log n); its merge step is reused in "count inversions" and "reverse pairs" (Step 3 hard).
- **Quick sort**: fastest in practice; its partition idea comes up in "kth largest" problems.
- In real Go code: `sort.Ints(a)` or `slices.Sort(a)`.

## Divide and conquer
Merge sort and quick sort both:
1. **Divide** the problem into smaller parts,
2. **Conquer** each part with recursion,
3. **Combine** the results.

Splitting in half again and again gives **log n levels**; each level does O(n) work → **O(n log n)**.
