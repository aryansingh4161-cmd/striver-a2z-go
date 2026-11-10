// Package trie is a basic lowercase trie, reused in Step 17.
package trie

// Node is one trie node; the letters live on the edges.
type Node struct {
	child [26]*Node
	end   bool
}

// Trie stores lowercase words.
type Trie struct{ root *Node }

// New returns an empty trie.
func New() *Trie { return &Trie{root: &Node{}} }

// Insert adds word to the trie. O(len(word)).
func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if cur.child[c] == nil {
			cur.child[c] = &Node{}
		}
		cur = cur.child[c]
	}
	cur.end = true
}

// walk follows s from the root and returns the last node, or nil if the path breaks.
func (t *Trie) walk(s string) *Node {
	cur := t.root
	for i := 0; i < len(s) && cur != nil; i++ {
		cur = cur.child[s[i]-'a']
	}
	return cur
}

// Search reports whether word was inserted. O(len(word)).
func (t *Trie) Search(word string) bool {
	n := t.walk(word)
	return n != nil && n.end
}

// StartsWith reports whether any inserted word starts with prefix. O(len(prefix)).
func (t *Trie) StartsWith(prefix string) bool { return t.walk(prefix) != nil }
