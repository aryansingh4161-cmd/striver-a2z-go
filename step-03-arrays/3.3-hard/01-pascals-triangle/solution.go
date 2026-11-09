package pascalstriangle

// Element returns the value at row r, column c (both 1-based) = C(r-1, c-1).
// Time O(c), space O(1).
func Element(r, c int) int {
	n, k := r-1, c-1
	res := 1
	for i := 0; i < k; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

// Row returns the r-th row (1-based). Time O(r).
func Row(r int) []int {
	row := []int{1}
	val := 1
	for c := 1; c < r; c++ {
		val = val * (r - c) / c
		row = append(row, val)
	}
	return row
}

// Generate returns the first n rows. Time O(n^2).
func Generate(n int) [][]int {
	res := make([][]int, 0, n)
	for r := 1; r <= n; r++ {
		res = append(res, Row(r))
	}
	return res
}
