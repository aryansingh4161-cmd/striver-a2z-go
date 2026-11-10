# Online Stock Span

**Problem (LeetCode 901):** Prices arrive one day at a time. For each day, return the **span**: the number of consecutive days up to and including today with price ≤ today's price.

Prices `100 80 60 70 60 75 85` → spans `1 1 1 2 1 4 6`

## Approach 1: Look back each day
**O(n) per day**, O(n²) total.

## Approach 2: Monotonic stack of (price, span) (optimal)
This is "**previous greater element**". Keep a stack of `(price, span)` with prices **decreasing**.
On a new price: start with `span = 1`; while the top's price `<= price`, pop it and **add its span** (those days are all covered). Push `(price, span)`.
- **Time:** O(1) amortised per day, **Space:** O(n)

## Dry run
| price | popped | span | stack after |
|---|---|---|---|
| 100 | – | 1 | (100,1) |
| 80 | – | 1 | (100,1)(80,1) |
| 60 | – | 1 | …(60,1) |
| 70 | (60,1) | 2 | (100,1)(80,1)(70,2) |
| 60 | – | 1 | …(60,1) |
| 75 | (60,1)(70,2) | 4 | (100,1)(80,1)(75,4) |
| 85 | (75,4)(80,1) | 6 | (100,1)(85,6) |
