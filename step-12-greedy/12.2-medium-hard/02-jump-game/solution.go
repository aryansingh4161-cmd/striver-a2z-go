package jumpgame

// CanJump reports whether the last index is reachable. O(n), O(1).
func CanJump(a []int) bool {
	farthest := 0
	for i, x := range a {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+x)
	}
	return true
}
