# Step 9: Stack and Queue

> Status: ✅ theory and all 30 problems done

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Stack using array](9.1-learning/01-stack-using-array/) | top index; slice version |
| 2 | [Queue using array](9.1-learning/02-queue-using-array/) | circular indexes with % cap |
| 3 | [Stack using queue](9.1-learning/03-stack-using-queue/) | rotate after push |
| 4 | [Queue using stacks](9.1-learning/04-queue-using-stack/) | in/out stacks, amortised O(1) |
| 5 | [Stack using linked list](9.1-learning/05-stack-using-linked-list/) | push/pop at head |
| 6 | [Queue using linked list](9.1-learning/06-queue-using-linked-list/) | head + tail pointers |
| 7 | [Valid parentheses](9.1-learning/07-valid-parentheses/) | stack of openers |
| 8 | [Min stack](9.1-learning/08-min-stack/) | (value, minSoFar) pairs / 2x−min trick |

### 9.2 Prefix, infix, postfix conversions
All six live in one package: [9.2-expression-conversions](9.2-expression-conversions/).

| # | Problem | Key idea |
|---|---|---|
| 1 | [Infix → postfix](9.2-expression-conversions/) | shunting-yard with precedence |
| 2 | [Prefix → infix](9.2-expression-conversions/) | right to left, bracket a op b |
| 3 | [Prefix → postfix](9.2-expression-conversions/) | right to left, a b op |
| 4 | [Postfix → prefix](9.2-expression-conversions/) | left to right, op a b |
| 5 | [Postfix → infix](9.2-expression-conversions/) | left to right, bracket a op b |
| 6 | [Infix → prefix](9.2-expression-conversions/) | reverse, shunting-yard, reverse |

### 9.3 Monotonic stack
| # | Problem | Key idea |
|---|---|---|
| 1 | [Next greater element](9.3-monotonic-stack/01-next-greater-element/) | decreasing stack from the right |
| 2 | [Next greater element II (circular)](9.3-monotonic-stack/02-next-greater-element-ii/) | loop 2n with i % n |
| 3 | [Next / previous smaller element](9.3-monotonic-stack/03-next-smaller-element/) | increasing stack |
| 4 | [Number of greater elements to the right](9.3-monotonic-stack/04-count-greater-to-right/) | scan per query / merge sort counts |
| 5 | [Trapping rain water](9.3-monotonic-stack/05-trapping-rain-water/) | two pointers / stack layers |
| 6 | [Sum of subarray minimums](9.3-monotonic-stack/06-sum-subarray-minimums/) | contribution = a[i]·left·right |
| 7 | [Asteroid collision](9.3-monotonic-stack/07-asteroid-collision/) | stack of survivors |
| 8 | [Sum of subarray ranges](9.3-monotonic-stack/08-sum-subarray-ranges/) | Σmax − Σmin |
| 9 | [Remove K digits](9.3-monotonic-stack/09-remove-k-digits/) | pop bigger digits greedily |
| 10 | [Largest rectangle in histogram](9.3-monotonic-stack/10-largest-rectangle-histogram/) | one-pass increasing stack |
| 11 | [Maximal rectangle](9.3-monotonic-stack/11-maximal-rectangle/) | each row is a histogram |

### 9.4 Implementation
| # | Problem | Key idea |
|---|---|---|
| 1 | [Sliding window maximum](9.4-implementation/01-sliding-window-maximum/) | monotonic deque of indexes |
| 2 | [Stock span](9.4-implementation/02-stock-span/) | (price, span) stack |
| 3 | [Celebrity problem](9.4-implementation/03-celebrity-problem/) | eliminate one per question |
| 4 | [LRU cache](9.4-implementation/04-lru-cache/) | map + doubly linked list |
| 5 | [LFU cache](9.4-implementation/05-lfu-cache/) | freq → list, track minFreq |
