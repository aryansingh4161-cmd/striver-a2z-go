package distancenearestzero

// UpdateMatrix returns each cell's distance to the nearest 0. Multi-source BFS. O(n*m).
func UpdateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	dist := make([][]int, n)
	var q [][2]int
	for r := range dist {
		dist[r] = make([]int, m)
		for c := range dist[r] {
			if mat[r][c] == 0 {
				q = append(q, [2]int{r, c})
			} else {
				dist[r][c] = -1 // not reached yet
			}
		}
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, d := range dirs {
			r, c := cur[0]+d[0], cur[1]+d[1]
			if r >= 0 && c >= 0 && r < n && c < m && dist[r][c] == -1 {
				dist[r][c] = dist[cur[0]][cur[1]] + 1
				q = append(q, [2]int{r, c})
			}
		}
	}
	return dist
}
