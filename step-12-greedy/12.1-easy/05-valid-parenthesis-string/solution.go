package validparenthesisstring

// CheckValidString tracks the range [lo, hi] of possible open counts. O(n), O(1).
func CheckValidString(s string) bool {
	lo, hi := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			lo++
			hi++
		case ')':
			lo--
			hi--
		default: // '*'
			lo--
			hi++
		}
		if hi < 0 {
			return false
		}
		lo = max(lo, 0)
	}
	return lo == 0
}

// CheckValidStringBrute tries all three meanings of every '*'. O(3^n).
func CheckValidStringBrute(s string) bool {
	var f func(i, open int) bool
	f = func(i, open int) bool {
		if open < 0 {
			return false
		}
		if i == len(s) {
			return open == 0
		}
		switch s[i] {
		case '(':
			return f(i+1, open+1)
		case ')':
			return f(i+1, open-1)
		}
		return f(i+1, open+1) || f(i+1, open-1) || f(i+1, open)
	}
	return f(0, 0)
}
