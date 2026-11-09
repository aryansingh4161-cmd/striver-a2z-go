# 1.6 Basic Hashing

**Problem it solves:** "How many times does X appear?" asked many times.
Scanning the whole array for every question costs O(n) each. **Hashing** means counting once and then answering each question in O(1).

## 1. Number hashing with an array

If values are small (say 0 to 10⁶), use them as indexes:

```go
hash := make([]int, maxVal+1)
for _, x := range arr {
    hash[x]++              // pre-compute: O(n)
}
fmt.Println(hash[5])       // fetch: O(1)
```

Limits: an array of ~10⁷ ints inside a function is fine in Go (Go puts big slices on the heap), but values like 10⁹ can't be indexes. Then use a **map**.

## 2. Character hashing

| Characters | Array size | Index |
|---|---|---|
| only `a`–`z` | 26 | `c - 'a'` |
| any ASCII | 256 | `int(c)` |

```go
var freq [26]int
for i := 0; i < len(s); i++ {
    freq[s[i]-'a']++
}
```

## 3. Map hashing (any values)

```go
freq := map[int]int{}
for _, x := range arr {
    freq[x]++
}
```

| Structure | Average | Worst case | Order |
|---|---|---|---|
| Go `map` | O(1) | O(n) (many collisions, very rare) | random |
| C++ `map` | O(log n) | O(log n) | sorted |
| C++ `unordered_map` | O(1) | O(n) | random |

## 4. How a hash table works inside

1. A **hash function** turns a key into a bucket number. Simplest one is the **division method**: `bucket = key % tableSize`.
2. Two keys can land in the same bucket (**collision**), e.g. 14 and 24 with size 10.
3. **Chaining:** each bucket keeps a small list of keys. Lookups only search that bucket's list.
4. If all keys collide, one bucket holds everything → O(n) lookup. Good hash functions make this extremely unlikely.

## Problems
| # | Problem |
|---|---|
| 1 | [Count frequency of each element](01-count-frequency/) |
| 2 | [Highest / lowest frequency element](02-highest-lowest-frequency/) |
