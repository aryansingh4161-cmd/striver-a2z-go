# Find the Repeating and Missing Numbers

**Problem:** The array has n numbers from `1..n`. One number appears twice (A) and one is missing (B). Return `[A, B]`.

`[3 1 2 5 3]` → `[3, 4]`

## Approach 1: Count each (brute)
For each value 1..n, count it. **O(n²)**.

## Approach 2: Hash array (better)
`count[x]++`; the value with 2 is A, the one with 0 is B. **Time O(n), Space O(n)**.

## Approach 3: Maths (optimal)
Let `x` = repeating, `y` = missing.
- `S  = sum(a)  − n(n+1)/2`  = **x − y**
- `S2 = sum(a²) − n(n+1)(2n+1)/6` = x² − y² = (x − y)(x + y) → **x + y = S2 / S**

Two equations → `x = (S + S2/S) / 2`, `y = x − S`.
- **Time:** O(n), **Space:** O(1)
- Overflow: `sum(a²)` gets large. Go's 64-bit `int` handles n up to about 10⁶ safely.

## Approach 4: XOR (optimal, no overflow)
1. `xr` = XOR of all elements and all of 1..n = **x ^ y** (everything else cancels).
2. x and y differ in at least one bit: take the **lowest set bit** `bit = xr & -xr`.
3. Split all numbers (array + 1..n) into two groups by that bit; XOR each group → one gives `x`, the other `y`.
4. Check which one appears in the array to tell which is the repeating one.

- **Time:** O(n), **Space:** O(1)

## Dry run (maths) `[3 1 2 5 3]`, n = 5
`S = 14 − 15 = −1` (x − y). `S2 = 48 − 55 = −7`. `x + y = −7 / −1 = 7`. `x = (−1 + 7)/2 = 3`, `y = 3 − (−1) = 4`. ✓
