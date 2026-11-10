# Implement a Queue Using an Array (circular)

**Queue = FIFO.** Operations: `Push(x)` (enqueue at the back), `Pop()` (dequeue from the front), `Front()`, `Size()`.

## The problem with a plain array
If you always pop from index 0 and shift everything left, pop is O(n). If you just move a `front` index forward, the space before it is wasted forever.

## Circular array (optimal)
Keep `front`, `rear` and `count`. Indexes **wrap around** with `% capacity`, so freed slots get reused.
- **Push:** if `count == cap` → full; `rear = (rear + 1) % cap`, `arr[rear] = x`, `count++`
- **Pop:** if `count == 0` → empty; `x = arr[front]`, `front = (front + 1) % cap`, `count--`
- All **O(1)**

```
cap 4, after push 1,2,3 then pop, pop, push 4,5:
index:  0  1  2  3
value:  5  .  3  4      front = 2, rear = 0 (wrapped)
```

In everyday Go, a slice with `q = q[1:]` is fine for BFS-style queues.
