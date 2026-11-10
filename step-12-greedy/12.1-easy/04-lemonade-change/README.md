# Lemonade Change

**Problem (LeetCode 860):** Lemonade costs 5. Customers pay with 5, 10 or 20, in order. You start with no change. Can you give everyone correct change?

`[5 5 5 10 20]` → true
`[5 5 10 10 20]` → false

## Greedy rule
Track how many 5s and 10s you have.
- pay 5 → `five++`
- pay 10 → need one 5: `five--`, `ten++`
- pay 20 → need 15: **prefer `10 + 5`** (keeps 5s, which are more flexible); otherwise `5 + 5 + 5`.
If you can't, return false.
- **Time:** O(n), **Space:** O(1)

Why prefer 10+5? A 5 can help a 10-payer or a 20-payer; a 10 can only help a 20-payer. Spend the less useful note first.
