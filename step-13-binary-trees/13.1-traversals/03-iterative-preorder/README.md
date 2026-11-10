# Iterative Preorder (with a stack)

Recursion uses the call stack; here we use our own.
```
stack = [root]
while stack not empty:
    node = pop
    visit node
    push node.Right   // pushed first …
    push node.Left    // … so Left is popped first
```
- **Time:** O(n), **Space:** O(h)

Pushing **right before left** is the whole trick: a stack is LIFO.
