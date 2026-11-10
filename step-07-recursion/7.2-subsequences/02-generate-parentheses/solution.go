package generateparentheses

// Generate returns all valid strings of n pairs of parentheses.
func Generate(n int) []string {
	res := []string{}
	buf := make([]byte, 0, 2*n)
	var f func(open, close int)
	f = func(open, close int) {
		if len(buf) == 2*n {
			res = append(res, string(buf))
			return
		}
		if open < n {
			buf = append(buf, '(')
			f(open+1, close)
			buf = buf[:len(buf)-1]
		}
		if close < open {
			buf = append(buf, ')')
			f(open, close+1)
			buf = buf[:len(buf)-1]
		}
	}
	f(0, 0)
	return res
}
