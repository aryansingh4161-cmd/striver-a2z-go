# 1.3 Go Collections (the "STL" lecture in Go)

Striver's C++ STL lecture covers containers like `vector`, `map`, `set`, `stack`, `queue`, `priority_queue`.
Go has fewer built-ins, but you can build all of them from **slices**, **maps** and `container/heap`.

The full guide with code is in [docs/go-for-dsa.md](../../docs/go-for-dsa.md). This page is the quick summary.

| C++ STL | Go | Main operations and cost |
|---|---|---|
| `vector` | slice `[]T` | append O(1) amortised, index O(1), insert in middle O(n) |
| `pair` | `[2]int` or struct | |
| `list` (doubly linked) | `container/list` (rarely needed) | |
| `stack` | slice | push/pop at end O(1) |
| `queue` | slice (`q = q[1:]`) | push O(1), pop O(1) |
| `deque` | slice | |
| `priority_queue` | `container/heap` | push/pop O(log n), top O(1) |
| `set` (sorted) | sorted slice / map + sort keys | Go has no balanced tree |
| `unordered_set` | `map[T]bool` | add/check O(1) average |
| `map` (sorted) | map + `sort` the keys when you need order | |
| `unordered_map` | `map[K]V` | O(1) average |
| `multimap` / `multiset` | `map[T]int` (value = count) | |
| `sort()` | `sort.Ints`, `sort.Slice`, `slices.Sort` | O(n log n) |
| `next_permutation` | write it yourself (Step 3 medium problem) | |
| `__builtin_popcount` | `bits.OnesCount(uint(x))` | |
| `*max_element` | `slices.Max(a)` | O(n) |

## Stack example

```go
st := []int{}
st = append(st, 1, 2, 3) // push
top := st[len(st)-1]     // peek → 3
st = st[:len(st)-1]      // pop
empty := len(st) == 0
```

## Queue example

```go
q := []int{}
q = append(q, 1, 2)  // push back
front := q[0]        // peek → 1
q = q[1:]            // pop front
```

## Sorted unique keys from a map

```go
keys := make([]int, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
```
