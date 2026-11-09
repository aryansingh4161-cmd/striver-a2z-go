# Second Largest Element

**Problem:** Return the second largest **distinct** element, or −1 if it doesn't exist.

| Input | Output |
|---|---|
| `[1 2 4 7 7 5]` | 5 |
| `[10 10 10]` | −1 |

## Approach 1: Sort (brute)
Sort, then walk from the end to find the first element smaller than the largest.
**Time O(n log n)**.

## Approach 2: Two passes (better)
Pass 1: find `largest`. Pass 2: find the biggest element that is `< largest`.
**Time O(2n) = O(n)**.

## Approach 3: One pass (optimal)
Keep `largest` and `second`:
```
for x in a:
    if x > largest:            second = largest; largest = x
    else if x < largest and x > second: second = x
```
- **Time:** O(n), **Space:** O(1)

## Dry run `[1 2 4 7 7 5]`
| x | largest | second |
|---|---|---|
| 1 | 1 | −∞ |
| 2 | 2 | 1 |
| 4 | 4 | 2 |
| 7 | 7 | 4 |
| 7 | 7 | 4 (equal, ignored) |
| 5 | 7 | **5** |

## Second smallest
Same idea with the signs flipped.

Use `math.MinInt` as "−∞" so negative arrays work.
