# Step 17: Tries

> Status: ✅ theory and all 6 problems done

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

| # | Problem | Key idea |
|---|---|---|
| 1 | [Implement trie (insert, search, startsWith)](01-implement-trie/) | 26 children + end flag |
| 2 | [Trie II (count words equal to / starting with, erase)](02-trie-ii/) | endCnt and preCnt in every node |
| 3 | [Longest word with all prefixes present](03-longest-complete-string/) | every node on the path must end a word |
| 4 | [Number of distinct substrings](04-distinct-substrings/) | insert all suffixes, count new nodes |
| 5 | [Maximum XOR of two numbers in an array](05-maximum-xor/) | bit trie, go for the opposite bit |
| 6 | [Maximum XOR with an element from an array (queries)](06-maximum-xor-queries/) | offline: sort queries by limit, grow the trie |
