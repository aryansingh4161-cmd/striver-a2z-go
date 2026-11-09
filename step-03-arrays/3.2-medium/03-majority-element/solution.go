package majorityelement

// MajorityMap counts with a map. Returns -1 if there is no majority.
// Time O(n), space O(n).
func MajorityMap(a []int) int {
	freq := map[int]int{}
	for _, x := range a {
		freq[x]++
		if freq[x] > len(a)/2 {
			return x
		}
	}
	return -1
}

// Majority uses Moore's Voting Algorithm, then verifies the candidate.
// Returns -1 if there is no majority. Time O(n), space O(1).
func Majority(a []int) int {
	candidate, count := 0, 0
	for _, x := range a {
		switch {
		case count == 0:
			candidate, count = x, 1
		case x == candidate:
			count++
		default:
			count--
		}
	}
	// verify
	c := 0
	for _, x := range a {
		if x == candidate {
			c++
		}
	}
	if c > len(a)/2 {
		return candidate
	}
	return -1
}
