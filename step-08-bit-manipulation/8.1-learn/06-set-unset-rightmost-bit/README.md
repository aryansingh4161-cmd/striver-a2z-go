# Set the Rightmost Unset Bit / Unset the Rightmost Set Bit

## Unset (clear) the rightmost set bit
`n & (n−1)`
`12 = 1100` → `1100 & 1011 = 1000` = 8

## Set the rightmost unset bit
`n | (n+1)`
Adding 1 turns the lowest 0 into 1 (and the 1s below it into 0s); OR-ing with n puts those 1s back.
`11 = 1011` → `1011 | 1100 = 1111` = 15

## Isolate the rightmost set bit
`n & −n`
`12 = 1100` → `−12 = ...0100` → `1100 & 0100 = 0100` = 4
(Used in "two numbers that appear once" in 8.2.)

All **O(1)**.
