# Add Two Numbers

**Problem (LeetCode 2):** Two numbers stored as lists with digits in **reverse order** (head = ones digit). Return their sum as a list in the same format.

`2 → 4 → 3` (342) + `5 → 6 → 4` (465) → `7 → 0 → 8` (807)

## Approach (optimal)
Exactly like adding on paper from the right, and the lists already start from the right.
```
dummy, carry = new node, 0
while a != nil or b != nil or carry != 0:
    sum = carry + (a's digit or 0) + (b's digit or 0)
    append node(sum % 10)
    carry = sum / 10
    advance a and b
```
Keeping `carry != 0` in the loop condition handles a final carry (e.g. 5 + 5 = `0 → 1`).
- **Time:** O(max(n, m)), **Space:** O(max(n, m)) for the answer
