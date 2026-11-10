# Implement a Trie (insert, search, startsWith)

**Problem:** Build a data structure with:
- `Insert(word)`
- `Search(word)`: is this exact word stored?
- `StartsWith(prefix)`: is any stored word starting with this prefix?

Words use lowercase `a`–`z`.

## Why not a hash set?
A set answers `Search` in O(L), but `StartsWith` would need to check every word: O(N · L). A trie answers both in **O(L)** (L = length of the word/prefix), no matter how many words are stored.

## The node
```go
type Node struct {
    child [26]*Node  // child[c] = next node for letter 'a'+c
    end   bool       // a word ends exactly here
}
```
The root is an empty node. The **path** from the root spells the word; the letters live on the edges.

## The three operations
```
Insert(w):   walk letters; create missing children; mark end on the last node
Search(w):   walk letters; if a child is missing → false; at the end return node.end
StartsWith(p): walk letters; if a child is missing → false; otherwise → true
```
`Search` and `StartsWith` differ **only** in the last line: the prefix doesn't need `end`.

## Dry run
Insert "apple" → path a-p-p-l-e, mark e.
- `Search("apple")` → reach e, `end = true` → **true**
- `Search("app")` → reach the second p, `end = false` → **false**
- `StartsWith("app")` → path exists → **true**
- Insert "app" → marks the second p as end → `Search("app")` → **true**

| | Time | Space |
|---|---|---|
| each operation | O(L) | Insert adds at most L nodes |
| total memory | | O(total letters × 26 pointers) |

This package is reused by the next problems.
