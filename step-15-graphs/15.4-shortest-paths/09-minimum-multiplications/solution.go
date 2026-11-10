package minimummultiplications

const mod = 100000

// MinSteps returns the fewest multiplications (mod 100000) from start to end,
// or -1. BFS over the 100000 possible values.
func MinSteps(arr []int, start, end int) int {
	if start == end {
		return 0
	}
	dist := make([]int, mod)
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	queue := []int{start}
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, a := range arr {
			y := x * a % mod
			if dist[y] != -1 {
				continue
			}
			dist[y] = dist[x] + 1
			if y == end {
				return dist[y]
			}
			queue = append(queue, y)
		}
	}
	return -1
}
