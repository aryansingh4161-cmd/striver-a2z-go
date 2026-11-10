package citysmallestneighbours

const inf = 1 << 60

// FindCity returns the city with the fewest others within threshold
// (ties → largest index). Floyd-Warshall, O(n³).
func FindCity(n int, edges [][3]int, threshold int) int {
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			if i != j {
				d[i][j] = inf
			}
		}
	}
	for _, e := range edges {
		d[e[0]][e[1]] = min(d[e[0]][e[1]], e[2])
		d[e[1]][e[0]] = min(d[e[1]][e[0]], e[2])
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j] // inf = 2^60 so inf+inf can't overflow
				}
			}
		}
	}
	best, city := n+1, -1
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i != j && d[i][j] <= threshold {
				cnt++
			}
		}
		if cnt <= best {
			best, city = cnt, i
		}
	}
	return city
}
