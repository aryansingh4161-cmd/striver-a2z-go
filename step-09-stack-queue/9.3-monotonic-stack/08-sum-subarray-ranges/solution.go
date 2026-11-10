package sumsubarrayranges

// RangesBrute keeps running max and min for every start. O(n^2).
func RangesBrute(a []int) int {
	total := 0
	for i := range a {
		lo, hi := a[i], a[i]
		for j := i; j < len(a); j++ {
			lo, hi = min(lo, a[j]), max(hi, a[j])
			total += hi - lo
		}
	}
	return total
}

// Ranges returns (sum of subarray maxima) - (sum of subarray minima). O(n).
func Ranges(a []int) int {
	return sumMax(a) - sumMin(a)
}

// sumMin: a[i] is the minimum of (i - prevSmaller) * (nextSmallerOrEqual - i) subarrays.
func sumMin(a []int) int {
	n := len(a)
	prev, next := make([]int, n), make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && a[st[len(st)-1]] >= a[i] {
			st = st[:len(st)-1]
		}
		prev[i] = -1
		if len(st) > 0 {
			prev[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[st[len(st)-1]] > a[i] {
			st = st[:len(st)-1]
		}
		next[i] = n
		if len(st) > 0 {
			next[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	total := 0
	for i := range a {
		total += a[i] * (i - prev[i]) * (next[i] - i)
	}
	return total
}

// sumMax: same as sumMin with the comparisons flipped (greater instead of smaller).
func sumMax(a []int) int {
	n := len(a)
	prev, next := make([]int, n), make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && a[st[len(st)-1]] <= a[i] {
			st = st[:len(st)-1]
		}
		prev[i] = -1
		if len(st) > 0 {
			prev[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[st[len(st)-1]] < a[i] {
			st = st[:len(st)-1]
		}
		next[i] = n
		if len(st) > 0 {
			next[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	total := 0
	for i := range a {
		total += a[i] * (i - prev[i]) * (next[i] - i)
	}
	return total
}
