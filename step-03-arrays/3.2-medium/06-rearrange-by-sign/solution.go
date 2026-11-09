package rearrangebysign

// RearrangeEqual alternates positive/negative when counts are equal.
// Time O(n), space O(n).
func RearrangeEqual(a []int) []int {
	res := make([]int, len(a))
	p, q := 0, 1
	for _, x := range a {
		if x > 0 {
			res[p] = x
			p += 2
		} else {
			res[q] = x
			q += 2
		}
	}
	return res
}

// RearrangeAny alternates as long as both signs remain, then appends the rest.
// Time O(n), space O(n).
func RearrangeAny(a []int) []int {
	var pos, neg []int
	for _, x := range a {
		if x > 0 {
			pos = append(pos, x)
		} else {
			neg = append(neg, x)
		}
	}
	res := make([]int, 0, len(a))
	k := min(len(pos), len(neg))
	for i := 0; i < k; i++ {
		res = append(res, pos[i], neg[i])
	}
	res = append(res, pos[k:]...)
	res = append(res, neg[k:]...)
	return res
}
