package majorityelementn3

import "sort"

// MajorityMap counts with a map. Time O(n), space O(n).
func MajorityMap(a []int) []int {
	freq := map[int]int{}
	res := []int{}
	for _, x := range a {
		freq[x]++
		if freq[x] == len(a)/3+1 { // just crossed n/3
			res = append(res, x)
		}
	}
	sort.Ints(res)
	return res
}

// Majority uses extended Moore's voting with two candidates.
// Result is sorted. Time O(n), space O(1).
func Majority(a []int) []int {
	cand1, cand2, c1, c2 := 0, 0, 0, 0
	for _, x := range a {
		switch {
		case c1 == 0 && x != cand2:
			cand1, c1 = x, 1
		case c2 == 0 && x != cand1:
			cand2, c2 = x, 1
		case x == cand1:
			c1++
		case x == cand2:
			c2++
		default:
			c1--
			c2--
		}
	}
	// verify
	c1, c2 = 0, 0
	for _, x := range a {
		if x == cand1 {
			c1++
		} else if x == cand2 {
			c2++
		}
	}
	res := []int{}
	if c1 > len(a)/3 {
		res = append(res, cand1)
	}
	if c2 > len(a)/3 && cand2 != cand1 {
		res = append(res, cand2)
	}
	sort.Ints(res)
	return res
}
