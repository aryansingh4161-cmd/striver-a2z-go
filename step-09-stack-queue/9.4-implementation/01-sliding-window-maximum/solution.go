package slidingwindowmaximum

// MaxBrute scans each window. O(n*k).
func MaxBrute(a []int, k int) []int {
	res := []int{}
	for i := 0; i+k <= len(a); i++ {
		m := a[i]
		for j := i; j < i+k; j++ {
			m = max(m, a[j])
		}
		res = append(res, m)
	}
	return res
}

// MaxSlidingWindow uses a monotonic deque of indexes. Time O(n), space O(k).
func MaxSlidingWindow(a []int, k int) []int {
	res := []int{}
	dq := []int{} // indexes; values decreasing from front to back
	for i := range a {
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:] // front left the window
		}
		for len(dq) > 0 && a[dq[len(dq)-1]] <= a[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if i >= k-1 {
			res = append(res, a[dq[0]])
		}
	}
	return res
}
