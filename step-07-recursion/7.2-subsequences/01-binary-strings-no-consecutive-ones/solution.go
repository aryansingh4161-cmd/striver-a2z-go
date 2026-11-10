package binarystrings

// Generate returns all length-n binary strings without "11", in sorted order.
func Generate(n int) []string {
	res := []string{}
	buf := make([]byte, 0, n)
	var f func()
	f = func() {
		if len(buf) == n {
			res = append(res, string(buf))
			return
		}
		buf = append(buf, '0')
		f()
		buf = buf[:len(buf)-1]
		if len(buf) == 0 || buf[len(buf)-1] == '0' {
			buf = append(buf, '1')
			f()
			buf = buf[:len(buf)-1]
		}
	}
	f()
	return res
}
