package spiralmatrix

// Spiral returns the elements of m in spiral order. Time O(nm).
func Spiral(m [][]int) []int {
	if len(m) == 0 {
		return []int{}
	}
	res := make([]int, 0, len(m)*len(m[0]))
	top, bottom, left, right := 0, len(m)-1, 0, len(m[0])-1
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ { // →
			res = append(res, m[top][j])
		}
		top++
		for i := top; i <= bottom; i++ { // ↓
			res = append(res, m[i][right])
		}
		right--
		if top <= bottom {
			for j := right; j >= left; j-- { // ←
				res = append(res, m[bottom][j])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- { // ↑
				res = append(res, m[i][left])
			}
			left++
		}
	}
	return res
}
