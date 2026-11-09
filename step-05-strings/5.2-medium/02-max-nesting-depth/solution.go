package maxnestingdepth

// MaxDepth returns the deepest nesting of parentheses. Time O(n), space O(1).
func MaxDepth(s string) int {
	depth, best := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
			best = max(best, depth)
		case ')':
			depth--
		}
	}
	return best
}
