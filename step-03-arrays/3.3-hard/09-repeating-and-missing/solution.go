package repeatingandmissing

// FindHash uses a count array. Returns [repeating, missing]. Time O(n), space O(n).
func FindHash(a []int) [2]int {
	n := len(a)
	count := make([]int, n+1)
	for _, x := range a {
		count[x]++
	}
	var res [2]int
	for v := 1; v <= n; v++ {
		switch count[v] {
		case 2:
			res[0] = v
		case 0:
			res[1] = v
		}
	}
	return res
}

// FindMath solves two equations from sum and sum of squares. Time O(n), space O(1).
func FindMath(a []int) [2]int {
	n := len(a)
	s, s2 := 0, 0
	for _, x := range a {
		s += x
		s2 += x * x
	}
	s -= n * (n + 1) / 2              // x - y
	s2 -= n * (n + 1) * (2*n + 1) / 6 // x^2 - y^2
	sumXY := s2 / s                   // x + y
	x := (s + sumXY) / 2
	y := x - s
	return [2]int{x, y}
}

// FindXOR uses the XOR split trick. Time O(n), space O(1).
func FindXOR(a []int) [2]int {
	xr := 0
	for i, x := range a {
		xr ^= x ^ (i + 1)
	}
	bit := xr & -xr // lowest set bit where x and y differ
	zero, one := 0, 0
	for i, x := range a {
		for _, v := range []int{x, i + 1} {
			if v&bit != 0 {
				one ^= v
			} else {
				zero ^= v
			}
		}
	}
	for _, x := range a { // find out which group's value is the repeating one
		if x == zero {
			return [2]int{zero, one}
		}
	}
	return [2]int{one, zero}
}
