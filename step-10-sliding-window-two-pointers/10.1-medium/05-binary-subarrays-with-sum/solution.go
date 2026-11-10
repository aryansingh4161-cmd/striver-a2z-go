package binarysubarrayswithsum

// CountPrefix uses prefix sums with a map. Time O(n), space O(n).
func CountPrefix(a []int, goal int) int {
	seen := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, x := range a {
		sum += x
		ans += seen[sum-goal]
		seen[sum]++
	}
	return ans
}

// Count returns atMost(goal) - atMost(goal-1). Time O(n), space O(1).
func Count(a []int, goal int) int {
	return atMost(a, goal) - atMost(a, goal-1)
}

func atMost(a []int, g int) int {
	if g < 0 {
		return 0
	}
	l, sum, count := 0, 0, 0
	for r, x := range a {
		sum += x
		for sum > g {
			sum -= a[l]
			l++
		}
		count += r - l + 1
	}
	return count
}
