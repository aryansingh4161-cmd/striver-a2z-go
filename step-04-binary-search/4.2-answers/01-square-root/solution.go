package squareroot

// SqrtLinear tries 1, 2, 3, ... Time O(sqrt(n)).
func SqrtLinear(n int) int {
	i := 0
	for (i+1)*(i+1) <= n {
		i++
	}
	return i
}

// Sqrt returns floor(sqrt(n)) by binary searching the answer. Time O(log n).
func Sqrt(n int) int {
	lo, hi, ans := 0, min(n, 3_037_000_499), 0 // hi*hi fits in int64
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid <= n {
			ans = mid // works; try bigger
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}
