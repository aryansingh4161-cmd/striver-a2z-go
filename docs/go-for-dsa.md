# Go for DSA: Everything You Need

Striver's videos use C++ and Java. This page shows the Go version of every tool those videos use, so you can translate on the fly.

> Keep this page open while watching lectures.

---

## 1. Translation table (C++ / Java → Go)

| C++ / Java | Go | Notes |
|---|---|---|
| `vector<int>` / `ArrayList<Integer>` | `[]int` (slice) | `append`, `len` |
| `int arr[n]` | `arr := make([]int, n)` | all values start at 0 |
| `vector<vector<int>>` | `[][]int` | make each row yourself (see below) |
| `unordered_map<K,V>` / `HashMap` | `map[K]V` | random iteration order |
| `map<K,V>` (sorted) / `TreeMap` | sort the keys, or sorted slice + `sort.Search` | Go has no ordered map |
| `set` / `HashSet` | `map[int]bool` or `map[int]struct{}` | no built-in set |
| `stack` | slice: push = `append`, pop = `s = s[:len(s)-1]` | |
| `queue` | slice: push = `append`, pop = `q = q[1:]` | |
| `deque` | slice (both ends) | |
| `priority_queue` | `container/heap` | see section 8 |
| `pair<int,int>` | `[2]int` or a small struct | |
| `INT_MAX` / `Integer.MAX_VALUE` | `math.MaxInt` | `math.MinInt` for minimum |
| `sort(v.begin(), v.end())` | `sort.Ints(v)` or `slices.Sort(v)` | |
| `swap(a, b)` | `a, b = b, a` | |
| `string` (mutable in C++) | `string` is **immutable** | convert to `[]byte` to change |
| `long long` | `int` (64-bit on all normal machines) | `int64` if you want to be explicit |

---

## 2. Slices (the most important thing)

```go
a := []int{5, 3, 8}        // literal
b := make([]int, 5)        // [0 0 0 0 0]
c := make([]int, 0, 10)    // length 0, capacity 10

a = append(a, 7)           // add at end
n := len(a)                // length
part := a[1:3]             // elements at index 1 and 2 (end is excluded)

for i, v := range a {      // index and value
    fmt.Println(i, v)
}
```

### 2D slices (grid / matrix)

```go
rows, cols := 3, 4
grid := make([][]int, rows)
for i := range grid {
    grid[i] = make([]int, cols)
}
```

### The trap: slices share memory

```go
a := []int{1, 2, 3}
b := a          // b points to the SAME data
b[0] = 99       // a[0] is now 99 too!

c := append([]int{}, a...) // real copy
d := make([]int, len(a))
copy(d, a)                 // also a real copy
```

This matters a lot in **recursion / backtracking**: when you save a path into your answer list, save a **copy**.

---

## 3. Maps

```go
freq := make(map[int]int)
freq[5]++                  // missing keys start at 0

if v, ok := freq[7]; ok {  // check if key exists
    fmt.Println("found", v)
}
delete(freq, 5)

for key, val := range freq { // order is RANDOM every run
    fmt.Println(key, val)
}
```

Set:

```go
seen := map[int]bool{}
seen[4] = true
if seen[4] { /* present */ }
```

---

## 4. Strings, bytes and runes

```go
s := "hello"
s[0]          // byte 'h' (type byte = uint8)
len(s)        // number of BYTES

b := []byte(s) // make it changeable
b[0] = 'H'
s = string(b)  // "Hello"

for i, r := range s { // r is a rune (a full character)
    fmt.Println(i, string(r))
}

r := []rune("héllo") // use for non-English text
```

Building a string fast:

```go
var sb strings.Builder
sb.WriteString("abc")
sb.WriteByte('d')
res := sb.String()
```

Useful packages: `strings` (`Split`, `Fields`, `Contains`, `Repeat`, `ToLower`), `strconv` (`Itoa`, `Atoi`).

Char tricks used in DSA:

```go
c := byte('7')
digit := int(c - '0')      // 7
idx := int('d' - 'a')      // 3 (position in alphabet)
```

---

## 5. Structs and pointers (linked lists, trees)

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

type TreeNode struct {
    Val         int
    Left, Right *TreeNode
}

head := &ListNode{Val: 1}
head.Next = &ListNode{Val: 2}
```

- Pass a **pointer** (`*T`) when the function must change the original.
- `nil` means "no node".

---

## 6. Sorting

```go
sort.Ints(a)                         // ascending
sort.Sort(sort.Reverse(sort.IntSlice(a))) // descending

sort.Slice(people, func(i, j int) bool { // custom order
    return people[i].age < people[j].age
})

slices.Sort(a)                       // Go 1.21+
idx, found := slices.BinarySearch(a, 8)
```

---

## 7. Recursion with closures

Helper functions that need access to outer variables:

```go
func subsets(nums []int) [][]int {
    var res [][]int
    var path []int

    var dfs func(i int)        // declare first so it can call itself
    dfs = func(i int) {
        if i == len(nums) {
            res = append(res, append([]int{}, path...)) // save a COPY
            return
        }
        path = append(path, nums[i])
        dfs(i + 1)
        path = path[:len(path)-1]
        dfs(i + 1)
    }
    dfs(0)
    return res
}
```

---

## 8. Heap (priority queue)

Go makes you implement 5 methods once. Copy this template:

```go
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // > for max-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// usage
h := &MinHeap{5, 2, 8}
heap.Init(h)
heap.Push(h, 1)
smallest := heap.Pop(h).(int) // 1
top := (*h)[0]                // peek
```

---

## 9. Small utilities

```go
math.MaxInt, math.MinInt   // "infinity"
min(a, b), max(a, b)       // built in since Go 1.21
math.Abs(float64(x))       // abs returns float64; for ints write your own:

func abs(x int) int {
    if x < 0 { return -x }
    return x
}
```

Integer division drops the decimal: `7 / 2 == 3`. Remainder: `7 % 2 == 1`.

---

## 10. Fast input (only for Codeforces-style stdin problems)

LeetCode doesn't need this.

```go
reader := bufio.NewReader(os.Stdin)
writer := bufio.NewWriter(os.Stdout)
defer writer.Flush()

var n int
fmt.Fscan(reader, &n)
fmt.Fprintln(writer, n)
```

---

## 11. Running code in this repo

```bash
go test ./...                           # run every test
go test ./step-03-arrays/...            # one step
go test -v ./step-03-arrays/3.1-easy/01-largest-element/  # one problem, verbose
```
