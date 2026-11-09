package countfrequency

// Pair holds a value and how many times it appears.
type Pair struct {
	Value, Count int
}

// FrequencyBrute uses a visited array. Time O(n^2).
func FrequencyBrute(a []int) []Pair {
	visited := make([]bool, len(a))
	var res []Pair
	for i := range a {
		if visited[i] {
			continue
		}
		count := 1
		for j := i + 1; j < len(a); j++ {
			if a[j] == a[i] {
				visited[j] = true
				count++
			}
		}
		res = append(res, Pair{a[i], count})
	}
	return res
}

// Frequency uses a map, keeping results in first-seen order. Time O(n).
func Frequency(a []int) []Pair {
	freq := map[int]int{}
	var order []int
	for _, x := range a {
		if freq[x] == 0 {
			order = append(order, x) // first time we see x
		}
		freq[x]++
	}
	res := make([]Pair, 0, len(order))
	for _, x := range order {
		res = append(res, Pair{x, freq[x]})
	}
	return res
}
