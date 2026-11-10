package distinctsubstrings

// Brute puts every substring in a set. O(n³).
func Brute(s string) int {
	set := map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			set[s[i:j]] = true
		}
	}
	return len(set)
}

type node struct{ child [26]*node }

// CountDistinct inserts every suffix into a trie and counts new nodes. O(n²).
func CountDistinct(s string) int {
	root := &node{}
	count := 0
	for i := 0; i < len(s); i++ {
		cur := root
		for j := i; j < len(s); j++ {
			c := s[j] - 'a'
			if cur.child[c] == nil {
				cur.child[c] = &node{}
				count++ // a substring we have never seen
			}
			cur = cur.child[c]
		}
	}
	return count
}
