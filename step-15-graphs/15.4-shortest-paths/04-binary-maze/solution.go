package binarymaze

// ShortestPath returns the fewest 4-directional steps from src to dst
// through cells equal to 1, or -1. O(rows · cols).
func ShortestPath(grid [][]int, src, dst [2]int) int {
	rows, cols := len(grid), len(grid[0])
	if grid[src[0]][src[1]] == 0 || grid[dst[0]][dst[1]] == 0 {
		return -1
	}
	if src == dst {
		return 0
	}
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[src[0]][src[1]] = 0
	queue := [][2]int{src}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			r, k := c[0]+d[0], c[1]+d[1]
			if r < 0 || r >= rows || k < 0 || k >= cols || grid[r][k] == 0 || dist[r][k] != -1 {
				continue
			}
			dist[r][k] = dist[c[0]][c[1]] + 1
			if r == dst[0] && k == dst[1] {
				return dist[r][k]
			}
			queue = append(queue, [2]int{r, k})
		}
	}
	return -1
}
