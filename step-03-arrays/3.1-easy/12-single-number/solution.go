package singlenumber

// SingleMap counts with a map. Time O(n), space O(n).
func SingleMap(a []int) int {
	freq := map[int]int{}
	for _, x := range a {
		freq[x]++
	}
	for x, c := range freq {
		if c == 1 {
			return x
		}
	}
	return 0
}

// SingleXOR XORs everything; pairs cancel. Time O(n), space O(1).
func SingleXOR(a []int) int {
	x := 0
	for _, v := range a {
		x ^= v
	}
	return x
}
