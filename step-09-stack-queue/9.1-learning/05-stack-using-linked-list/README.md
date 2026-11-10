# Implement a Stack Using a Linked List

Push and pop at the **head** of a singly linked list: both are O(1), and there's no fixed capacity.
```
push(x): node = {x, next: top}; top = node; size++
pop():   x = top.val; top = top.next; size--
```
- All operations **O(1)**, memory grows and shrinks with the stack.
