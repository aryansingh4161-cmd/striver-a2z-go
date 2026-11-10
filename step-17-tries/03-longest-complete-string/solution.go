package longestcompletestring

type node struct {
	child [26]*node
	end   bool
}

// Brute checks every prefix of every word in a set. O(N·L²).
func Brute(words []string) string {
	set := map[string]bool{}
	for _, w := range words {
		set[w] = true
	}
	best := ""
	for _, w := range words {
		ok := true
		for l := 1; l <= len(w); l++ {
			if !set[w[:l]] {
				ok = false
				break
			}
		}
		if ok && (len(w) > len(best) || (len(w) == len(best) && w < best)) {
			best = w
		}
	}
	if best == "" {
		return "None"
	}
	return best
}

// CompleteString uses a trie: every node on the word's path must end a word. O(N·L).
func CompleteString(words []string) string {
	root := &node{}
	for _, w := range words {
		cur := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if cur.child[c] == nil {
				cur.child[c] = &node{}
			}
			cur = cur.child[c]
		}
		cur.end = true
	}
	best := ""
	for _, w := range words {
		cur, ok := root, true
		for i := 0; i < len(w); i++ {
			cur = cur.child[w[i]-'a']
			if !cur.end {
				ok = false
				break
			}
		}
		if ok && (len(w) > len(best) || (len(w) == len(best) && w < best)) {
			best = w
		}
	}
	if best == "" {
		return "None"
	}
	return best
}
