package ratinmaze

// FindPaths returns all paths from (0,0) to (n-1,n-1), sorted, as DLRU strings.
func FindPaths(maze [][]int) []string {
	n := len(maze)
	res := []string{}
	if n == 0 || maze[0][0] == 0 || maze[n-1][n-1] == 0 {
		return res
	}
	dr := []int{1, 0, 0, -1}
	dc := []int{0, -1, 1, 0}
	name := "DLRU" // alphabetical order → sorted output
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	path := []byte{}
	var f func(r, c int)
	f = func(r, c int) {
		if r == n-1 && c == n-1 {
			res = append(res, string(path))
			return
		}
		visited[r][c] = true
		for k := 0; k < 4; k++ {
			nr, nc := r+dr[k], c+dc[k]
			if nr >= 0 && nc >= 0 && nr < n && nc < n && maze[nr][nc] == 1 && !visited[nr][nc] {
				path = append(path, name[k])
				f(nr, nc)
				path = path[:len(path)-1]
			}
		}
		visited[r][c] = false
	}
	f(0, 0)
	return res
}
