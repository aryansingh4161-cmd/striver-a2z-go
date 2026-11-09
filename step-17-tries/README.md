# Step 17: Tries

> Status: theory ready, problems coming ⏳

## What it is
A tree for storing strings where each edge is a character. Words that share a prefix share the path.

```
insert "app", "apple", "bat"

        (root)
        /    \
       a      b
       |      |
       p      a
       |      |
       p*     t*
       |
       l
       |
       e*          (* = a word ends here)
```

```go
type Node struct {
    Child [26]*Node
    End   bool // a word ends at this node
}

func (n *Node) Insert(w string) {
    cur := n
    for i := 0; i < len(w); i++ {
        c := w[i] - 'a'
        if cur.Child[c] == nil {
            cur.Child[c] = &Node{}
        }
        cur = cur.Child[c]
    }
    cur.End = true
}
```

| Operation | Time |
|---|---|
| insert / search / startsWith | O(length of word) |

## When to use
- Many queries like "does any word start with …?" (autocomplete).
- Counting distinct substrings (insert every suffix).
- **Maximum XOR:** a trie over **bits** (2 children: 0 and 1). For each number, greedily walk towards the opposite bit to maximise XOR.

## Problems
| # | Problem | Status |
|---|---|---|
| 1 | Implement trie (insert, search, startsWith) | ⏳ |
| 2 | Trie II (count words equal to / starting with, erase) | ⏳ |
| 3 | Longest word with all prefixes present | ⏳ |
| 4 | Number of distinct substrings | ⏳ |
| 5 | Maximum XOR of two numbers in an array | ⏳ |
| 6 | Maximum XOR with an element from an array (queries) | ⏳ |
