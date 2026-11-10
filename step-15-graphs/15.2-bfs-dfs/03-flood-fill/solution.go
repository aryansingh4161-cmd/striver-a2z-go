package floodfill

// FloodFill repaints the region containing (sr, sc). Returns a new image. O(n*m).
func FloodFill(image [][]int, sr, sc, color int) [][]int {
	img := make([][]int, len(image))
	for i := range image {
		img[i] = append([]int{}, image[i]...)
	}
	orig := img[sr][sc]
	if orig == color {
		return img
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(img) || c >= len(img[0]) || img[r][c] != orig {
			return
		}
		img[r][c] = color
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	dfs(sr, sc)
	return img
}
