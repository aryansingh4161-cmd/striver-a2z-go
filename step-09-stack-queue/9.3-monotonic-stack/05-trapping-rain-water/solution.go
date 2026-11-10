package trappingrainwater

// TrapPrefix uses prefix and suffix maxima. Time O(n), space O(n).
func TrapPrefix(h []int) int {
	n := len(h)
	if n == 0 {
		return 0
	}
	left, right := make([]int, n), make([]int, n)
	left[0] = h[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], h[i])
	}
	right[n-1] = h[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], h[i])
	}
	total := 0
	for i := range h {
		total += min(left[i], right[i]) - h[i]
	}
	return total
}

// Trap uses two pointers. Time O(n), space O(1).
func Trap(h []int) int {
	l, r := 0, len(h)-1
	leftMax, rightMax, total := 0, 0, 0
	for l < r {
		if h[l] <= h[r] {
			leftMax = max(leftMax, h[l])
			total += leftMax - h[l]
			l++
		} else {
			rightMax = max(rightMax, h[r])
			total += rightMax - h[r]
			r--
		}
	}
	return total
}

// TrapStack fills water layer by layer with a monotonic stack. Time O(n), space O(n).
func TrapStack(h []int) int {
	st := []int{} // indexes, heights decreasing
	total := 0
	for i := range h {
		for len(st) > 0 && h[i] > h[st[len(st)-1]] {
			mid := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			width := i - left - 1
			height := min(h[left], h[i]) - h[mid]
			total += width * height
		}
		st = append(st, i)
	}
	return total
}
