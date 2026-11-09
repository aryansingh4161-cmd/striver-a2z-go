package singleelementsortedarray

// SingleXOR XORs everything. Time O(n).
func SingleXOR(a []int) int {
	x := 0
	for _, v := range a {
		x ^= v
	}
	return x
}

// Single uses binary search on index parity. Time O(log n), space O(1).
func Single(a []int) int {
	n := len(a)
	if n == 1 {
		return a[0]
	}
	if a[0] != a[1] {
		return a[0]
	}
	if a[n-1] != a[n-2] {
		return a[n-1]
	}
	lo, hi := 1, n-2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] != a[mid-1] && a[mid] != a[mid+1] {
			return a[mid]
		}
		// in the left (correct-pairing) zone?
		if (mid%2 == 1 && a[mid] == a[mid-1]) || (mid%2 == 0 && a[mid] == a[mid+1]) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
