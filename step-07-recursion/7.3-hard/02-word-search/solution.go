package wordsearch

// Exist reports whether word can be traced in the grid. The grid is restored
// before returning. Time O(n*m*3^L), space O(L).
func Exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	n, m := len(board), len(board[0])
	var dfs func(r, c, k int) bool
	dfs = func(r, c, k int) bool {
		if r < 0 || c < 0 || r >= n || c >= m || board[r][c] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		saved := board[r][c]
		board[r][c] = '#' // mark as used
		found := dfs(r+1, c, k+1) || dfs(r-1, c, k+1) || dfs(r, c+1, k+1) || dfs(r, c-1, k+1)
		board[r][c] = saved // backtrack
		return found
	}
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
