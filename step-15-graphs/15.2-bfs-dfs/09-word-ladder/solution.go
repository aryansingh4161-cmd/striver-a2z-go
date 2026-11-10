package wordladder

// LadderLength returns the number of words in the shortest transformation, or 0.
// BFS over words. Time O(N * L * 26).
func LadderLength(begin, end string, wordList []string) int {
	words := map[string]bool{}
	for _, w := range wordList {
		words[w] = true
	}
	if !words[end] {
		return 0
	}
	delete(words, begin)
	q := []string{begin}
	steps := 1
	for len(q) > 0 {
		for size := len(q); size > 0; size-- {
			w := q[0]
			q = q[1:]
			if w == end {
				return steps
			}
			b := []byte(w)
			for i := range b {
				orig := b[i]
				for c := byte('a'); c <= 'z'; c++ {
					b[i] = c
					next := string(b)
					if words[next] {
						delete(words, next)
						q = append(q, next)
					}
				}
				b[i] = orig
			}
		}
		steps++
	}
	return 0
}
