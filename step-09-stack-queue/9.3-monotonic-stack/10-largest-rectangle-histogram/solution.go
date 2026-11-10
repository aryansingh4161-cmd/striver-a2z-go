package largestrectanglehistogram

// LargestBrute expands left and right from each bar. O(n^2).
func LargestBrute(h []int) int {
	best := 0
	for i := range h {
		l, r := i, i
		for l > 0 && h[l-1] >= h[i] {
			l--
		}
		for r < len(h)-1 && h[r+1] >= h[i] {
			r++
		}
		best = max(best, h[i]*(r-l+1))
	}
	return best
}

// Largest uses a one-pass increasing stack. Time O(n), space O(n).
func Largest(h []int) int {
	n := len(h)
	st := []int{}
	best := 0
	for i := 0; i <= n; i++ {
		cur := 0
		if i < n {
			cur = h[i]
		}
		for len(st) > 0 && h[st[len(st)-1]] >= cur {
			height := h[st[len(st)-1]]
			st = st[:len(st)-1]
			left := -1
			if len(st) > 0 {
				left = st[len(st)-1]
			}
			best = max(best, height*(i-left-1))
		}
		st = append(st, i)
	}
	return best
}
