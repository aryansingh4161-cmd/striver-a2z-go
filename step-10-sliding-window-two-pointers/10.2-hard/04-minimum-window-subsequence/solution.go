package minimumwindowsubsequence

// MinWindow returns the shortest substring of s1 that has s2 as a subsequence.
// Forward/backward two-pointer scan. Time O(n*m), space O(1).
func MinWindow(s1, s2 string) string {
	n, m := len(s1), len(s2)
	if m == 0 {
		return ""
	}
	bestStart, bestLen := -1, n+1
	i := 0
	for i < n {
		// forward: match all of s2 starting from i
		j := 0
		for i < n {
			if s1[i] == s2[j] {
				j++
				if j == m {
					break
				}
			}
			i++
		}
		if i == n {
			break
		}
		end := i
		// backward: tightest start for this end
		j = m - 1
		for j >= 0 {
			if s1[i] == s2[j] {
				j--
			}
			i--
		}
		start := i + 1
		if end-start+1 < bestLen {
			bestStart, bestLen = start, end-start+1
		}
		i = start + 1
	}
	if bestStart == -1 {
		return ""
	}
	return s1[bestStart : bestStart+bestLen]
}

// MinWindowBrute tries every start and walks forward. O(n^2).
func MinWindowBrute(s1, s2 string) string {
	best := ""
	for i := 0; i < len(s1); i++ {
		j := 0
		for k := i; k < len(s1) && len(s2) > 0; k++ {
			if s1[k] == s2[j] {
				j++
				if j == len(s2) {
					if best == "" || k-i+1 < len(best) {
						best = s1[i : k+1]
					}
					break
				}
			}
		}
	}
	return best
}
