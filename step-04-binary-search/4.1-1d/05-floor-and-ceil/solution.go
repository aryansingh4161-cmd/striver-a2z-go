package floorandceil

// Floor returns the largest element <= x, or -1. Time O(log n).
func Floor(a []int, x int) int {
	lo, hi, ans := 0, len(a)-1, -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] <= x {
			ans = a[mid] // fits; try to find a bigger one on the right
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

// Ceil returns the smallest element >= x, or -1. Time O(log n).
func Ceil(a []int, x int) int {
	lo, hi, ans := 0, len(a)-1, -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] >= x {
			ans = a[mid] // fits; try to find a smaller one on the left
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
