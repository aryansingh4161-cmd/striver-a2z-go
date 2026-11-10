package longeststringchain

import "sort"

// isPred reports whether inserting one letter into a gives b.
func isPred(a, b string) bool {
	if len(b) != len(a)+1 {
		return false
	}
	i := 0
	for j := 0; j < len(b); j++ {
		if i < len(a) && a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}

// LongestChain is LIS over words sorted by length. O(n²·L).
func LongestChain(words []string) int {
	w := append([]string(nil), words...)
	sort.Slice(w, func(i, j int) bool { return len(w[i]) < len(w[j]) })
	dp := make([]int, len(w))
	best := 0
	for i := range w {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isPred(w[j], w[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		best = max(best, dp[i])
	}
	return best
}

// LongestChainMap tries deleting each letter and looks the shorter word up. O(n·L²).
func LongestChainMap(words []string) int {
	w := append([]string(nil), words...)
	sort.Slice(w, func(i, j int) bool { return len(w[i]) < len(w[j]) })
	best, memo := 0, map[string]int{}
	for _, word := range w {
		cur := 1
		for k := 0; k < len(word); k++ {
			if v, ok := memo[word[:k]+word[k+1:]]; ok {
				cur = max(cur, v+1)
			}
		}
		memo[word] = max(memo[word], cur)
		best = max(best, cur)
	}
	return best
}
