package leftrotatebyone

// RotateLeftOne shifts a left by one place in place. Time O(n), space O(1).
func RotateLeftOne(a []int) {
	if len(a) < 2 {
		return
	}
	first := a[0]
	for i := 1; i < len(a); i++ {
		a[i-1] = a[i]
	}
	a[len(a)-1] = first
}
