package trieii

type node struct {
	child          [26]*node
	endCnt, preCnt int
}

// Trie supports duplicates, counting and erase. Every operation is O(L).
type Trie struct{ root *node }

// New returns an empty trie.
func New() *Trie { return &Trie{root: &node{}} }

// Insert adds one copy of word.
func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if cur.child[c] == nil {
			cur.child[c] = &node{}
		}
		cur = cur.child[c]
		cur.preCnt++
	}
	cur.endCnt++
}

func (t *Trie) walk(s string) *node {
	cur := t.root
	for i := 0; i < len(s) && cur != nil; i++ {
		cur = cur.child[s[i]-'a']
	}
	return cur
}

// CountWordsEqualTo returns how many copies of word are stored.
func (t *Trie) CountWordsEqualTo(word string) int {
	if n := t.walk(word); n != nil {
		return n.endCnt
	}
	return 0
}

// CountWordsStartingWith returns how many stored words start with prefix (prefix must be non-empty).
func (t *Trie) CountWordsStartingWith(prefix string) int {
	if n := t.walk(prefix); n != nil {
		return n.preCnt
	}
	return 0
}

// Erase removes one copy of word, which must be present.
func (t *Trie) Erase(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		cur = cur.child[word[i]-'a']
		cur.preCnt--
	}
	cur.endCnt--
}
