package twosinglenumbers

// TwoSingles returns the two numbers that appear once (smaller first).
// Time O(n), space O(1).
func TwoSingles(a []int) (int, int) {
	xr := 0
	for _, v := range a {
		xr ^= v
	}
	bit := xr & -xr // a bit where the two answers differ
	x, y := 0, 0
	for _, v := range a {
		if v&bit != 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	if x > y {
		x, y = y, x
	}
	return x, y
}
