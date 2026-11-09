package leftrotatebyd

// RotateLeftTemp uses a temporary slice of size d. Time O(n), space O(d).
func RotateLeftTemp(a []int, d int) {
	n := len(a)
	if n == 0 {
		return
	}
	d %= n
	temp := append([]int{}, a[:d]...)
	copy(a, a[d:])
	copy(a[n-d:], temp)
}

// RotateLeft uses the reversal algorithm. Time O(n), space O(1).
func RotateLeft(a []int, d int) {
	n := len(a)
	if n == 0 {
		return
	}
	d %= n
	reverse(a, 0, d-1)
	reverse(a, d, n-1)
	reverse(a, 0, n-1)
}

// RotateRight rotates right by k (LeetCode 189). Time O(n), space O(1).
func RotateRight(a []int, k int) {
	n := len(a)
	if n == 0 {
		return
	}
	k %= n
	reverse(a, 0, n-1)
	reverse(a, 0, k-1)
	reverse(a, k, n-1)
}

func reverse(a []int, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
