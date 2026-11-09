package missingnumber

// MissingHash marks seen numbers. Time O(n), space O(n).
func MissingHash(a []int) int {
	n := len(a)
	seen := make([]bool, n+1)
	for _, x := range a {
		seen[x] = true
	}
	for v := 0; v <= n; v++ {
		if !seen[v] {
			return v
		}
	}
	return -1
}

// MissingSum uses expected sum minus actual sum. Time O(n), space O(1).
func MissingSum(a []int) int {
	n := len(a)
	sum := 0
	for _, x := range a {
		sum += x
	}
	return n*(n+1)/2 - sum
}

// MissingXOR cancels pairs with XOR. Time O(n), space O(1).
func MissingXOR(a []int) int {
	x := len(a) // include n itself
	for i, v := range a {
		x ^= i ^ v
	}
	return x
}
