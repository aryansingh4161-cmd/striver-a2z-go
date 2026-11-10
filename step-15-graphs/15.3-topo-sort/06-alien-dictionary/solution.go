package aliendictionary

// FindOrder returns an alphabet order of the first k letters that is
// consistent with the sorted word list, or "" if no order exists.
func FindOrder(words []string, k int) string {
	adj := make([][]int, k)
	indeg := make([]int, k)
	for i := 0; i+1 < len(words); i++ {
		a, b := words[i], words[i+1]
		n := min(len(a), len(b))
		found := false
		for j := 0; j < n; j++ {
			if a[j] != b[j] {
				u, v := int(a[j]-'a'), int(b[j]-'a')
				adj[u] = append(adj[u], v) // a[j] comes before b[j]
				indeg[v]++
				found = true
				break // only the first difference tells us anything
			}
		}
		if !found && len(a) > len(b) {
			return "" // "abc" before "ab" is impossible
		}
	}
	queue := []int{}
	for c := 0; c < k; c++ {
		if indeg[c] == 0 {
			queue = append(queue, c)
		}
	}
	order := make([]byte, 0, k)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, byte('a'+u))
		for _, v := range adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(order) < k {
		return "" // cycle in the rules
	}
	return string(order)
}
