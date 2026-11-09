package printnamentimes

// PrintName returns name repeated n times (instead of printing, so it can be tested).
func PrintName(name string, n int) []string {
	var out []string
	var f func(k int)
	f = func(k int) {
		if k == 0 { // base case
			return
		}
		out = append(out, name)
		f(k - 1)
	}
	f(n)
	return out
}
