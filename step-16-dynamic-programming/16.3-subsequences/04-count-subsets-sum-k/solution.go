package countsubsetssumk

const mod = 1_000_000_007

// Recursive counts subsets with sum k (handles zeros). O(2^n).
func Recursive(arr []int, k int) int {
	var f func(i, t int) int
	f = func(i, t int) int {
		if i == 0 {
			switch {
			case t == 0 && arr[0] == 0:
				return 2
			case t == 0 || t == arr[0]:
				return 1
			}
			return 0
		}
		res := f(i-1, t)
		if arr[i] <= t {
			res += f(i-1, t-arr[i])
		}
		return res % mod
	}
	return f(len(arr)-1, k)
}

// Memo caches f(i, t). O(n·k).
func Memo(arr []int, k int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, k+1)
		for t := range dp[i] {
			dp[i][t] = -1
		}
	}
	var f func(i, t int) int
	f = func(i, t int) int {
		if i == 0 {
			switch {
			case t == 0 && arr[0] == 0:
				return 2
			case t == 0 || t == arr[0]:
				return 1
			}
			return 0
		}
		if dp[i][t] != -1 {
			return dp[i][t]
		}
		res := f(i-1, t)
		if arr[i] <= t {
			res += f(i-1, t-arr[i])
		}
		dp[i][t] = res % mod
		return dp[i][t]
	}
	return f(len(arr)-1, k)
}

// Tab is the full table. O(n·k) time and space.
func Tab(arr []int, k int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	if arr[0] == 0 {
		dp[0][0] = 2
	} else {
		dp[0][0] = 1
		if arr[0] <= k {
			dp[0][arr[0]] = 1
		}
	}
	for i := 1; i < n; i++ {
		for t := 0; t <= k; t++ {
			dp[i][t] = dp[i-1][t]
			if arr[i] <= t {
				dp[i][t] = (dp[i][t] + dp[i-1][t-arr[i]]) % mod
			}
		}
	}
	return dp[n-1][k]
}

// Optimal uses one array, t from high to low. O(k) space.
func Optimal(arr []int, k int) int {
	ways := make([]int, k+1)
	ways[0] = 1
	for _, x := range arr {
		for t := k; t >= x; t-- {
			ways[t] = (ways[t] + ways[t-x]) % mod
		}
	}
	return ways[k]
}
