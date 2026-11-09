package timesarrayrotated

// Rotations returns how many times a sorted array of distinct values was
// rotated right = the index of its minimum. Time O(log n).
func Rotations(a []int) int {
	lo, hi := 0, len(a)-1
	minIdx := 0
	for lo <= hi {
		if a[lo] <= a[hi] {
			if a[lo] < a[minIdx] {
				minIdx = lo
			}
			break
		}
		mid := lo + (hi-lo)/2
		if a[lo] <= a[mid] {
			if a[lo] < a[minIdx] {
				minIdx = lo
			}
			lo = mid + 1
		} else {
			if a[mid] < a[minIdx] {
				minIdx = mid
			}
			hi = mid - 1
		}
	}
	return minIdx
}
