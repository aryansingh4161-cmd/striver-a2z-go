package handofstraights

import "sort"

// IsNStraightHand reports whether hand can be split into groups of
// groupSize consecutive cards. Time O(n log n).
func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := map[int]int{}
	for _, c := range hand {
		count[c]++
	}
	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		c := count[v]
		if c == 0 {
			continue
		}
		for x := v; x < v+groupSize; x++ {
			if count[x] < c {
				return false
			}
			count[x] -= c
		}
	}
	return true
}
