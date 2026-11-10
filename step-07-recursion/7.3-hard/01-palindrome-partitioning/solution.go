package palindromepartitioning

// Partition returns all ways to split s into palindromic pieces.
func Partition(s string) [][]string {
	res := [][]string{}
	path := []string{}
	var f func(start int)
	f = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if isPal(s, start, i) {
				path = append(path, s[start:i+1])
				f(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	f(0)
	return res
}

func isPal(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
