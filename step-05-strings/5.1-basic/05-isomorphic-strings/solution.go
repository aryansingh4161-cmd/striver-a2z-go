package isomorphicstrings

// IsIsomorphicMaps checks both mapping directions with maps. Time O(n).
func IsIsomorphicMaps(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	st, ts := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if x, ok := st[a]; ok && x != b {
			return false
		}
		if y, ok := ts[b]; ok && y != a {
			return false
		}
		st[a], ts[b] = b, a
	}
	return true
}

// IsIsomorphic compares last-seen positions. Time O(n), space O(1).
func IsIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var lastS, lastT [256]int
	for i := 0; i < len(s); i++ {
		if lastS[s[i]] != lastT[t[i]] {
			return false
		}
		lastS[s[i]], lastT[t[i]] = i+1, i+1
	}
	return true
}
