package highestlowestfrequency

// HighLow returns the most frequent and least frequent elements of a (non-empty).
// On ties, the element that appears first in a wins. Time O(n), space O(k).
func HighLow(a []int) (most, least int) {
	freq := map[int]int{}
	var order []int
	for _, x := range a {
		if freq[x] == 0 {
			order = append(order, x)
		}
		freq[x]++
	}
	maxC, minC := 0, len(a)+1
	for _, x := range order {
		c := freq[x]
		if c > maxC {
			maxC, most = c, x
		}
		if c < minC {
			minC, least = c, x
		}
	}
	return most, least
}
