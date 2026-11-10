# Implement a Queue Using a Linked List

Keep **two** pointers: `head` (front, where we pop) and `tail` (back, where we push).
```
push(x): node = {x}; if tail == nil: head = tail = node
                     else: tail.next = node; tail = node
pop():   x = head.val; head = head.next; if head == nil: tail = nil
```
The last line matters: after removing the only element, `tail` must become nil too.
- All operations **O(1)**
