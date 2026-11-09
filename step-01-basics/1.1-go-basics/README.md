# 1.1 Go Basics + Time Complexity

Just enough Go to start DSA. Every snippet runs inside `func main()`.

## 1. Program skeleton

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, DSA")
}
```

Run with `go run main.go`.

## 2. Input and output

```go
var n int
fmt.Scan(&n)               // read one integer
var name string
fmt.Scan(&name)            // reads one word

fmt.Println("n =", n)       // prints with spaces + newline
fmt.Printf("%d %s\n", n, name) // formatted
```

| Verb | Prints |
|---|---|
| `%d` | integer |
| `%s` | string |
| `%f`, `%.2f` | float, 2 decimals |
| `%v` | anything (slices, maps, structs) |
| `%c` | a character from a byte/rune |

## 3. Data types

| Type | Example | Notes |
|---|---|---|
| `int` | `42` | 64-bit on normal machines, up to about 9.2 × 10^18 |
| `int32`, `int64` | | explicit size |
| `float64` | `3.14` | decimals |
| `bool` | `true` | |
| `string` | `"abc"` | immutable |
| `byte` | `'a'` | one ASCII character (alias of `uint8`) |
| `rune` | `'é'` | one Unicode character (alias of `int32`) |

```go
x := 10          // short declaration (type inferred)
var y int = 20   // explicit
const MOD = 1_000_000_007
```

Go never converts types for you: `float64(x) / 2`, `int(3.9) == 3`.

## 4. Conditions

```go
if age >= 18 {
    fmt.Println("adult")
} else if age >= 13 {
    fmt.Println("teen")
} else {
    fmt.Println("child")
}
```

Switch (no `break` needed, Go stops after a matching case):

```go
switch day {
case 1:
    fmt.Println("Monday")
case 6, 7:
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

## 5. Loops (Go only has `for`)

```go
for i := 0; i < n; i++ { }     // classic
for i < n { i++ }               // works like while
for { break }                   // infinite loop, exit with break
for i, v := range arr { }       // over a slice
```

## 6. Arrays and strings

```go
arr := [5]int{1, 2, 3}    // fixed-size array (rare in DSA)
s := []int{1, 2, 3}       // slice (use this)
str := "hello"
fmt.Println(len(str), str[1]) // 5 101  (str[1] is the byte 'e')
```

Full details: [Go for DSA guide](../../docs/go-for-dsa.md).

## 7. Functions

```go
func add(a, b int) int {
    return a + b
}

func divmod(a, b int) (int, int) { // multiple return values
    return a / b, a % b
}

q, r := divmod(17, 5) // 3 2
```

**Pass by value vs reference:**
- `int`, `string`, structs, arrays → a **copy** is passed.
- slices, maps, pointers → the function can change the **original** data.

```go
func setFirst(a []int) { a[0] = 100 } // changes caller's slice
func inc(p *int)       { *p++ }        // changes caller's int
```

---

## 8. Time complexity (Big-O)

Time complexity counts **how the number of steps grows** as the input size `n` grows. It is **not** the time in seconds (that depends on the machine).

### Rules
1. Look at the **worst case**.
2. Drop **constants**: `O(2n + 5)` → `O(n)`.
3. Keep only the **biggest term**: `O(n² + n)` → `O(n²)`.

### Examples

```go
for i := 0; i < n; i++ { }              // O(n)

for i := 0; i < n; i++ {               // O(n²)
    for j := 0; j < n; j++ { }
}

for i := 0; i < n; i++ {               // O(n²) too: n(n+1)/2 steps
    for j := 0; j <= i; j++ { }
}

for i := 1; i < n; i *= 2 { }          // O(log n): i doubles each time

for n > 0 { n /= 10 }                  // O(log₁₀ n): number of digits
```

### Growth order (fast → slow)

`O(1) < O(log n) < O(√n) < O(n) < O(n log n) < O(n²) < O(2ⁿ) < O(n!)`

### Will my code pass?

A judge runs about **10⁸ simple operations per second**.

| n up to | Needed complexity |
|---|---|
| 10 – 12 | O(n!) |
| 20 – 25 | O(2ⁿ) |
| 500 | O(n³) |
| 10⁴ | O(n²) |
| 10⁶ | O(n log n) |
| 10⁸ | O(n) |
| bigger | O(log n) or O(1) |

### Space complexity

Extra memory your algorithm uses **besides the input** (called *auxiliary space*). A new slice of size `n` → `O(n)`. Recursion depth `d` → `O(d)` stack space.

**Good habit:** don't change the input array unless the problem allows it.
