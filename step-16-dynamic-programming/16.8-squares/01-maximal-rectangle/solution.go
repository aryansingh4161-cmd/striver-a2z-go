package maximalrectangle

// largestInHistogram uses a stack of increasing bars. O(m).
func largestInHistogram(h []int) int {
	best := 0
	stack := []int{} // indexes with increasing heights
	for i := 0; i <= len(h); i++ {
		cur := 0
		if i < len(h) {
			cur = h[i]
		}
		for len(stack) > 0 && h[stack[len(stack)-1]] >= cur {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1 // nearest smaller bar on the left
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			best = max(best, h[top]*(i-left-1)) // i is the nearest smaller on the right
		}
		stack = append(stack, i)
	}
	return best
}

// MaximalRectangle builds a histogram per row. O(n·m).
func MaximalRectangle(mat [][]int) int {
	if len(mat) == 0 {
		return 0
	}
	height := make([]int, len(mat[0]))
	best := 0
	for _, row := range mat {
		for j, v := range row {
			if v == 1 {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		best = max(best, largestInHistogram(height))
	}
	return best
}

// Brute checks every rectangle. O(n²·m²·n·m), for tests only.
func Brute(mat [][]int) int {
	n, m, best := len(mat), len(mat[0]), 0
	for r1 := 0; r1 < n; r1++ {
		for c1 := 0; c1 < m; c1++ {
			for r2 := r1; r2 < n; r2++ {
				for c2 := c1; c2 < m; c2++ {
					ok := true
					for r := r1; r <= r2 && ok; r++ {
						for c := c1; c <= c2; c++ {
							if mat[r][c] == 0 {
								ok = false
								break
							}
						}
					}
					if ok {
						best = max(best, (r2-r1+1)*(c2-c1+1))
					}
				}
			}
		}
	}
	return best
}
