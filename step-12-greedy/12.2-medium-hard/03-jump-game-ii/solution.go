package jumpgameii

// MinJumps returns the fewest jumps to reach the last index (assumed reachable).
// Level-by-level greedy. Time O(n), space O(1).
func MinJumps(a []int) int {
	jumps, l, r := 0, 0, 0
	for r < len(a)-1 {
		farthest := 0
		for i := l; i <= r; i++ {
			farthest = max(farthest, i+a[i])
		}
		l, r = r+1, farthest
		jumps++
	}
	return jumps
}
