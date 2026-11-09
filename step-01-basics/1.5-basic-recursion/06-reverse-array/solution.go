package reversearray

// ReverseTwoPointers reverses a in place with recursion on (l, r).
func ReverseTwoPointers(a []int) {
	var f func(l, r int)
	f = func(l, r int) {
		if l >= r {
			return
		}
		a[l], a[r] = a[r], a[l]
		f(l+1, r-1)
	}
	f(0, len(a)-1)
}

// ReverseOnePointer reverses a in place, swapping i with n-1-i.
func ReverseOnePointer(a []int) {
	n := len(a)
	var f func(i int)
	f = func(i int) {
		if i >= n/2 {
			return
		}
		a[i], a[n-1-i] = a[n-1-i], a[i]
		f(i + 1)
	}
	f(0)
}
