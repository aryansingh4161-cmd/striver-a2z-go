package maximalrectangle

// MaximalRectangle returns the area of the largest all-ones rectangle.
// Each row becomes a histogram. Time O(n*m), space O(m).
func MaximalRectangle(mat [][]byte) int {
	if len(mat) == 0 {
		return 0
	}
	heights := make([]int, len(mat[0]))
	best := 0
	for _, row := range mat {
		for j, c := range row {
			if c == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		best = max(best, largestHistogram(heights))
	}
	return best
}

func largestHistogram(h []int) int {
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
