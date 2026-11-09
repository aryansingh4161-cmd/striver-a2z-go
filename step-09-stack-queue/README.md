# Step 9: Stack and Queue

> Status: theory ready, problems coming ⏳

## Stack: Last In, First Out (LIFO)
Like a pile of plates. Push on top, pop from top.
```go
st := []int{}
st = append(st, x)          // push
top := st[len(st)-1]        // peek
st = st[:len(st)-1]         // pop
```
Used for: matching brackets, undo, function calls, expression evaluation, monotonic stack problems.

## Queue: First In, First Out (FIFO)
Like a line at a counter.
```go
q := []int{}
q = append(q, x)            // push back
front := q[0]
q = q[1:]                   // pop front
```
Used for: BFS (Step 15), scheduling, sliding window maximum (with a deque).

## Monotonic stack (the most important pattern here)
Keep the stack in increasing (or decreasing) order. Before pushing `x`, pop everything that breaks the order. Each element is pushed and popped once → **O(n)** total.

**Next greater element** (scan from the right):
```go
res := make([]int, n)
st := []int{}
for i := n - 1; i >= 0; i-- {
    for len(st) > 0 && st[len(st)-1] <= a[i] {
        st = st[:len(st)-1]      // smaller ones can never be anyone's answer now
    }
    if len(st) == 0 { res[i] = -1 } else { res[i] = st[len(st)-1] }
    st = append(st, a[i])
}
```
Variants: next smaller, previous greater/smaller, circular array (loop `2n` times with `i % n`).
Uses: largest rectangle in histogram, trapping rain water, sum of subarray minimums, stock span, remove K digits.

## Expressions
- **Infix:** `a + b` (how humans write)
- **Postfix:** `a b +` (no brackets needed, easy to evaluate with a stack)
- **Prefix:** `+ a b`
Infix → postfix uses operator **precedence** with a stack (Shunting-yard).

## Design problems
- **Min stack:** push pairs `(value, minSoFar)`.
- **LRU cache:** hash map + doubly linked list → O(1) get and put.
- **LFU cache:** map of frequency → list.

## Problems

### 9.1 Learning
| # | Problem | Status |
|---|---|---|
| 1 | Implement stack using array | ⏳ |
| 2 | Implement queue using array | ⏳ |
| 3 | Implement stack using queue | ⏳ |
| 4 | Implement queue using stack | ⏳ |
| 5 | Implement stack using linked list | ⏳ |
| 6 | Implement queue using linked list | ⏳ |
| 7 | Valid parentheses | ⏳ |
| 8 | Min stack | ⏳ |

### 9.2 Prefix, infix, postfix conversions
| # | Problem | Status |
|---|---|---|
| 1 | Infix to postfix | ⏳ |
| 2 | Prefix to infix | ⏳ |
| 3 | Prefix to postfix | ⏳ |
| 4 | Postfix to prefix | ⏳ |
| 5 | Postfix to infix | ⏳ |
| 6 | Infix to prefix | ⏳ |

### 9.3 Monotonic stack / queue
| # | Problem | Status |
|---|---|---|
| 1 | Next greater element | ⏳ |
| 2 | Next greater element II (circular) | ⏳ |
| 3 | Next smaller element | ⏳ |
| 4 | Number of greater elements to the right | ⏳ |
| 5 | Trapping rain water | ⏳ |
| 6 | Sum of subarray minimums | ⏳ |
| 7 | Asteroid collision | ⏳ |
| 8 | Sum of subarray ranges | ⏳ |
| 9 | Remove K digits | ⏳ |
| 10 | Largest rectangle in histogram | ⏳ |
| 11 | Maximal rectangle | ⏳ |

### 9.4 Implementation
| # | Problem | Status |
|---|---|---|
| 1 | Sliding window maximum | ⏳ |
| 2 | Stock span | ⏳ |
| 3 | The celebrity problem | ⏳ |
| 4 | LRU cache | ⏳ |
| 5 | LFU cache | ⏳ |
