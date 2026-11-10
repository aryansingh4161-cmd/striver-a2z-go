package eventualsafestates

// SafeNodesDFS: a node is safe if DFS from it never meets a cycle. O(V + E).
func SafeNodesDFS(g [][]int) []int {
	n := len(g)
	visited, onPath := make([]bool, n), make([]bool, n)
	var dfs func(int) bool // true if a cycle is reachable
	dfs = func(u int) bool {
		visited[u], onPath[u] = true, true
		for _, v := range g[u] {
			if onPath[v] || (!visited[v] && dfs(v)) {
				return true // leave onPath[u] = true: u is unsafe
			}
		}
		onPath[u] = false
		return false
	}
	for s := 0; s < n; s++ {
		if !visited[s] {
			dfs(s)
		}
	}
	safe := []int{}
	for u := 0; u < n; u++ {
		if !onPath[u] {
			safe = append(safe, u)
		}
	}
	return safe
}

// SafeNodesKahn: reverse the edges and peel off nodes whose every outgoing
// edge already leads to a safe node. O(V + E).
func SafeNodesKahn(g [][]int) []int {
	n := len(g)
	rev := make([][]int, n)
	outdeg := make([]int, n)
	for u := range g {
		outdeg[u] = len(g[u])
		for _, v := range g[u] {
			rev[v] = append(rev[v], u)
		}
	}
	queue := []int{}
	for u := 0; u < n; u++ {
		if outdeg[u] == 0 {
			queue = append(queue, u)
		}
	}
	isSafe := make([]bool, n)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		isSafe[v] = true
		for _, u := range rev[v] {
			outdeg[u]--
			if outdeg[u] == 0 {
				queue = append(queue, u)
			}
		}
	}
	safe := []int{}
	for u := 0; u < n; u++ { // walking 0..n-1 gives sorted order
		if isSafe[u] {
			safe = append(safe, u)
		}
	}
	return safe
}
