package unionofsortedarrays

// Union returns the sorted distinct values from sorted slices a and b.
// Time O(n+m).
func Union(a, b []int) []int {
	res := []int{}
	add := func(x int) {
		if len(res) == 0 || res[len(res)-1] != x {
			res = append(res, x)
		}
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			add(a[i])
			i++
		} else {
			add(b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		add(a[i])
	}
	for ; j < len(b); j++ {
		add(b[j])
	}
	return res
}

// Intersection returns the common elements of sorted slices a and b,
// keeping repeats as many times as they appear in both. Time O(n+m).
func Intersection(a, b []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			j++
		default:
			res = append(res, a[i])
			i++
			j++
		}
	}
	return res
}
