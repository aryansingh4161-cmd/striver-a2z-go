# Implement a Stack Using a Queue

**Problem (LeetCode 225):** Build LIFO behaviour using only queue operations (push back, pop front, front, size).

## Approach: Rotate after every push (one queue)
After pushing `x` to the back, move every **older** element from the front to the back (`size − 1` times). Now `x` is at the front, so the front of the queue is always the top of the stack.
```
push(x):
    q.push(x)
    repeat size-1 times: q.push(q.pop())
pop(): return q.pop()
top(): return q.front()
```
- **Push:** O(n), **Pop / Top:** O(1)

## Picture: push 1, 2, 3
```
push 1: [1]
push 2: [1 2] → rotate once → [2 1]
push 3: [2 1 3] → rotate twice → [3 2 1]   front = 3 = top ✓
```
