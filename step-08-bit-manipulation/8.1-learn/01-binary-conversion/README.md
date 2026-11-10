# Decimal ↔ Binary Conversion

## Decimal → binary
Keep dividing by 2; the remainders, read **bottom-up**, are the bits.

13 → 13/2 = 6 r **1**, 6/2 = 3 r **0**, 3/2 = 1 r **1**, 1/2 = 0 r **1** → read up: `1101`
- **Time:** O(log n)

## Binary → decimal
Each bit is worth `2^position` (rightmost = position 0).
`1101` = 1·8 + 1·4 + 0·2 + 1·1 = 13

Walk left to right with `value = value*2 + bit` (same idea as building a decimal number with `*10`).
- **Time:** O(length)

## In Go
```go
strconv.FormatInt(13, 2)          // "1101"
v, _ := strconv.ParseInt("1101", 2, 64) // 13
fmt.Printf("%b", 13)              // 1101
```

## Negative numbers: two's complement
Computers store `−x` as **flip all bits of x, then add 1**.
In 8 bits: `5 = 00000101` → flip `11111010` → +1 → `11111011` = −5.
That's why `x & -x` and `^x` behave the way they do in later problems.
