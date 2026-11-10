# Trie II (count words, count prefixes, erase)

**Problem:** A trie that allows **duplicate** words, with:
- `Insert(word)`
- `CountWordsEqualTo(word)`: how many times was `word` inserted (and not erased)?
- `CountWordsStartingWith(prefix)`: how many stored words start with `prefix`?
- `Erase(word)`: remove one copy of `word` (it's guaranteed to exist).

## Store two counters in every node
```go
type Node struct {
    child  [26]*Node
    endCnt int  // words that END here
    preCnt int  // words that PASS THROUGH here (have this prefix)
}
```
- `Insert`: on the way down, `preCnt++` on every node you step into; at the end `endCnt++`.
- `CountWordsEqualTo(w)`: walk `w`; return `endCnt` of the last node (0 if the path breaks).
- `CountWordsStartingWith(p)`: walk `p`; return `preCnt` of the last node.
- `Erase(w)`: walk `w`, doing `preCnt--` on each node; at the end `endCnt--`.

## Dry run
Insert "apple", "apple", "apps", "apxl"
| node after | preCnt | endCnt |
|---|---|---|
| a | 4 | 0 |
| ap | 4 | 0 |
| app | 3 | 0 |
| apple (e) | 2 | 2 |
| apps (s) | 1 | 1 |
- `CountWordsEqualTo("apple")` = 2
- `CountWordsStartingWith("app")` = 3
- `Erase("apple")` → `CountWordsEqualTo("apple")` = 1, `CountWordsStartingWith("app")` = 2 ✅

All operations: **O(L)**.

(Nodes whose `preCnt` drops to 0 could be deleted to free memory; leaving them is simpler and still correct.)
