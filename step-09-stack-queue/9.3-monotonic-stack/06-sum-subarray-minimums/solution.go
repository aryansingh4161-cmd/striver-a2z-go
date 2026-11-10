package sumsubarrayminimums

const mod = 1_000_000_007

// SumMinsBrute checks every subarray. O(n^2).
func SumMinsBrute(a []int) int {
	total := 0
	for i := range a {
		m := a[i]
		for j := i; j < len(a); j++ {
			m = min(m, a[j])
			total = (total + m) % mod
		}
	}
	return total
}

// SumMins adds each element's contribution a[i] * left * right. O(n).
func SumMins(a []int) int {
	n := len(a)
	prevLess := make([]int, n) // index of previous strictly smaller, -1 if none
	nextLessEq := make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && a[st[len(st)-1]] >= a[i] {
			st = st[:len(st)-1]
		}
		prevLess[i] = -1
		if len(st) > 0 {
			prevLess[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[st[len(st)-1]] > a[i] {
			st = st[:len(st)-1]
		}
		nextLessEq[i] = n
		if len(st) > 0 {
			nextLessEq[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	total := 0
	for i := 0; i < n; i++ {
		left, right := i-prevLess[i], nextLessEq[i]-i
		total = (total + a[i]%mod*left%mod*right) % mod
	}
	return total
}
