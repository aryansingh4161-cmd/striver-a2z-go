# Step 8: Bit Manipulation (+ maths)

> Status: theory ready, problems coming ⏳

## Binary basics
- 13 in binary = `1101` = 8 + 4 + 0 + 1.
- **Decimal → binary:** keep dividing by 2, read remainders bottom-up.
- **Binary → decimal:** sum of `bit × 2^position`.
- Negative numbers use **two's complement**: `−x = ^x + 1` (flip all bits, add 1).

```go
fmt.Printf("%b\n", 13)               // 1101
n, _ := strconv.ParseInt("1101", 2, 64) // 13
```

## Operators in Go
| Operator | Go | Rule |
|---|---|---|
| AND | `a & b` | 1 only if both are 1 |
| OR | `a \| b` | 1 if either is 1 |
| XOR | `a ^ b` | 1 if they differ |
| NOT | `^a` | flips every bit (Go uses `^`, not `~`) |
| Left shift | `a << k` | × 2ᵏ |
| Right shift | `a >> k` | ÷ 2ᵏ (floor) |
| AND NOT | `a &^ b` | clears bits of a that are set in b (Go only) |

## Must-know tricks
| Task | Code |
|---|---|
| Is the i-th bit set? | `n>>i&1 == 1` or `n&(1<<i) != 0` |
| Set the i-th bit | `n \| (1<<i)` |
| Clear the i-th bit | `n &^ (1<<i)` |
| Toggle the i-th bit | `n ^ (1<<i)` |
| Clear the lowest set bit | `n & (n-1)` |
| Keep only the lowest set bit | `n & -n` |
| Is n a power of 2? | `n > 0 && n&(n-1) == 0` |
| Is n odd? | `n&1 == 1` |
| Count set bits | `bits.OnesCount(uint(n))` or loop `n &= n-1` |
| Swap without temp | `a ^= b; b ^= a; a ^= b` (Go: just `a, b = b, a`) |

## XOR facts
`x ^ x = 0`, `x ^ 0 = x`, order doesn't matter. Pairs cancel, which is how "single number" and "missing number" work.

XOR of 1..n follows a pattern by `n % 4`: `n, 1, n+1, 0`.

## Subsets with bitmasks
Numbers `0 .. 2ⁿ−1` are all the on/off patterns of n items. Bit `i` set → include item `i`.

## Maths part
- **Sieve of Eratosthenes:** all primes up to N in O(N log log N).
- **Smallest prime factor (SPF) sieve:** factorise many numbers fast.
- **Fast power (binary exponentiation):** `xⁿ` in O(log n) by squaring.

## Problems

### 8.1 Learn
| # | Problem | Status |
|---|---|---|
| 1 | Decimal ↔ binary, operators | ⏳ |
| 2 | Check if the i-th bit is set | ⏳ |
| 3 | Check odd or even | ⏳ |
| 4 | Check power of 2 | ⏳ |
| 5 | Count set bits | ⏳ |
| 6 | Set / unset the rightmost bit | ⏳ |
| 7 | Swap two numbers | ⏳ |
| 8 | Divide two integers without `*`, `/` | ⏳ |

### 8.2 Interview problems
| # | Problem | Status |
|---|---|---|
| 1 | Minimum bit flips to convert A to B | ⏳ |
| 2 | Single number (others appear twice) | ⏳ |
| 3 | Single number II (others appear thrice) | ⏳ |
| 4 | Power set using bits | ⏳ |
| 5 | XOR of numbers in range L..R | ⏳ |
| 6 | Two numbers that appear once (others twice) | ⏳ |

### 8.3 Advanced maths
| # | Problem | Status |
|---|---|---|
| 1 | Prime factors of a number | ⏳ |
| 2 | All divisors (√n) | ⏳ |
| 3 | Sieve of Eratosthenes | ⏳ |
| 4 | Prime factorisation with SPF sieve | ⏳ |
| 5 | Pow(x, n) with binary exponentiation | ⏳ |
