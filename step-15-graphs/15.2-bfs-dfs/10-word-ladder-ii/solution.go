package wordladderii

import "sort"

// FindLadders returns all shortest transformation sequences (sorted).
func FindLadders(begin, end string, wordList []string) [][]string {
	words := map[string]bool{}
	for _, w := range wordList {
		words[w] = true
	}
	res := [][]string{}
	if !words[end] {
		return res
	}
	delete(words, begin)
	depth := map[string]int{begin: 1}
	q := []string{begin}
	for len(q) > 0 && depth[end] == 0 {
		var usedThisLevel []string
		for size := len(q); size > 0; size-- {
			w := q[0]
			q = q[1:]
			for _, nb := range neighbours(w) {
				if words[nb] {
					if _, ok := depth[nb]; !ok {
						depth[nb] = depth[w] + 1
						q = append(q, nb)
					}
					usedThisLevel = append(usedThisLevel, nb)
				}
			}
		}
		for _, w := range usedThisLevel { // remove only after the whole level
			delete(words, w)
		}
	}
	if depth[end] == 0 {
		return res
	}
	path := []string{end}
	var dfs func(w string)
	dfs = func(w string) {
		if w == begin {
			p := make([]string, len(path))
			for i := range path {
				p[i] = path[len(path)-1-i]
			}
			res = append(res, p)
			return
		}
		for _, nb := range neighbours(w) {
			if d, ok := depth[nb]; ok && d+1 == depth[w] {
				path = append(path, nb)
				dfs(nb)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(end)
	sort.Slice(res, func(i, j int) bool {
		for k := range res[i] {
			if res[i][k] != res[j][k] {
				return res[i][k] < res[j][k]
			}
		}
		return false
	})
	return res
}

func neighbours(w string) []string {
	var out []string
	b := []byte(w)
	for i := range b {
		orig := b[i]
		for c := byte('a'); c <= 'z'; c++ {
			if c != orig {
				b[i] = c
				out = append(out, string(b))
			}
		}
		b[i] = orig
	}
	return out
}
