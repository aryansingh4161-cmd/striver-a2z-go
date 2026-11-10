# Step 8: Bit Manipulation (+ maths)

> Status: ✅ theory and all 19 problems done

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Decimal ↔ binary](8.1-learn/01-binary-conversion/) | divide by 2 / value*2 + bit |
| 2 | [Check / set / clear / toggle the i-th bit](8.1-learn/02-check-ith-bit/) | masks with 1 << i |
| 3 | [Odd or even](8.1-learn/03-odd-even/) | n & 1 |
| 4 | [Power of 2](8.1-learn/04-power-of-two/) | n & (n−1) == 0 |
| 5 | [Count set bits](8.1-learn/05-count-set-bits/) | Kernighan: n &= n−1 |
| 6 | [Set / unset the rightmost bit](8.1-learn/06-set-unset-rightmost-bit/) | n | (n+1), n & (n−1), n & −n |
| 7 | [Swap two numbers](8.1-learn/07-swap-numbers/) | XOR swap |
| 8 | [Divide without * / %](8.1-learn/08-divide-integers/) | subtract biggest divisor << k |

### 8.2 Interview problems
| # | Problem | Key idea |
|---|---|---|
| 1 | [Minimum bit flips A → B](8.2-interview/01-min-bit-flips/) | popcount(a ^ b) |
| 2 | [Single number (others twice)](8.2-interview/02-single-number/) | XOR all |
| 3 | [Single number II (others thrice)](8.2-interview/03-single-number-ii/) | bit counts mod 3 / ones-twos |
| 4 | [Power set with bits](8.2-interview/04-power-set-bits/) | mask 0 .. 2ⁿ−1 |
| 5 | [XOR of L..R](8.2-interview/05-xor-range/) | n % 4 pattern |
| 6 | [Two numbers appearing once](8.2-interview/06-two-single-numbers/) | split by lowest differing bit |

### 8.3 Advanced maths
| # | Problem | Key idea |
|---|---|---|
| 1 | [Prime factors](8.3-maths/01-prime-factors/) | divide out up to √n |
| 2 | [All divisors](8.3-maths/02-all-divisors/) | pairs up to √n |
| 3 | [Sieve of Eratosthenes](8.3-maths/03-sieve/) | cross out from p² |
| 4 | [Factorisation with SPF sieve](8.3-maths/04-spf-factorisation/) | divide by spf[x] repeatedly |
| 5 | [Pow(x, n) / modular power](8.3-maths/05-pow-binary-exp/) | square and multiply on bits |
