package ninjatraining

// Recursive tries every non-repeating schedule. O(3^n).
func Recursive(p [][3]int) int {
	var f func(day, last int) int
	f = func(day, last int) int {
		best := 0
		for act := 0; act < 3; act++ {
			if act == last {
				continue
			}
			v := p[day][act]
			if day > 0 {
				v += f(day-1, act)
			}
			best = max(best, v)
		}
		return best
	}
	return f(len(p)-1, 3)
}

// Memo caches f(day, last). O(n·4·3).
func Memo(p [][3]int) int {
	n := len(p)
	dp := make([][4]int, n)
	for i := range dp {
		dp[i] = [4]int{-1, -1, -1, -1}
	}
	var f func(day, last int) int
	f = func(day, last int) int {
		if dp[day][last] != -1 {
			return dp[day][last]
		}
		best := 0
		for act := 0; act < 3; act++ {
			if act == last {
				continue
			}
			v := p[day][act]
			if day > 0 {
				v += f(day-1, act)
			}
			best = max(best, v)
		}
		dp[day][last] = best
		return best
	}
	return f(n-1, 3)
}

// Tab fills dp[day][last] bottom-up. O(n) time, O(n) space.
func Tab(p [][3]int) int {
	n := len(p)
	dp := make([][4]int, n)
	for last := 0; last < 4; last++ {
		for act := 0; act < 3; act++ {
			if act != last {
				dp[0][last] = max(dp[0][last], p[0][act])
			}
		}
	}
	for day := 1; day < n; day++ {
		for last := 0; last < 4; last++ {
			for act := 0; act < 3; act++ {
				if act != last {
					dp[day][last] = max(dp[day][last], p[day][act]+dp[day-1][act])
				}
			}
		}
	}
	return dp[n-1][3]
}

// Optimal keeps only the previous day's row. O(n) time, O(1) space.
func Optimal(p [][3]int) int {
	var prev [4]int
	for last := 0; last < 4; last++ {
		for act := 0; act < 3; act++ {
			if act != last {
				prev[last] = max(prev[last], p[0][act])
			}
		}
	}
	for day := 1; day < len(p); day++ {
		var cur [4]int
		for last := 0; last < 4; last++ {
			for act := 0; act < 3; act++ {
				if act != last {
					cur[last] = max(cur[last], p[day][act]+prev[act])
				}
			}
		}
		prev = cur
	}
	return prev[3]
}
