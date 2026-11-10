package subsetsum

// Recursive: take or not take each element. O(2^n).
func Recursive(arr []int, target int) bool {
	var f func(i, t int) bool
	f = func(i, t int) bool {
		if t == 0 {
			return true
		}
		if i == 0 {
			return arr[0] == t
		}
		if f(i-1, t) {
			return true
		}
		return arr[i] <= t && f(i-1, t-arr[i])
	}
	return f(len(arr)-1, target)
}

// Memo caches f(i, t) (0 = unknown, 1 = false, 2 = true). O(n·target).
func Memo(arr []int, target int) bool {
	dp := make([][]int8, len(arr))
	for i := range dp {
		dp[i] = make([]int8, target+1)
	}
	var f func(i, t int) bool
	f = func(i, t int) bool {
		if t == 0 {
			return true
		}
		if i == 0 {
			return arr[0] == t
		}
		if dp[i][t] != 0 {
			return dp[i][t] == 2
		}
		res := f(i-1, t) || (arr[i] <= t && f(i-1, t-arr[i]))
		dp[i][t] = 1
		if res {
			dp[i][t] = 2
		}
		return res
	}
	return f(len(arr)-1, target)
}

// Tab is the full boolean table. O(n·target) time and space.
func Tab(arr []int, target int) bool {
	n := len(arr)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	if arr[0] <= target {
		dp[0][arr[0]] = true
	}
	for i := 1; i < n; i++ {
		for t := 1; t <= target; t++ {
			dp[i][t] = dp[i-1][t] || (arr[i] <= t && dp[i-1][t-arr[i]])
		}
	}
	return dp[n-1][target]
}

// Optimal uses one array, looping t from high to low. O(target) space.
func Optimal(arr []int, target int) bool {
	dp := make([]bool, target+1)
	dp[0] = true
	for _, x := range arr {
		for t := target; t >= x; t-- {
			dp[t] = dp[t] || dp[t-x] // dp[t-x] is still last round's value
		}
	}
	return dp[target]
}
