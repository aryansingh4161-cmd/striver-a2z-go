package lettercombinations

var keypad = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// LetterCombinations returns all strings the digits could spell.
func LetterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	buf := make([]byte, 0, len(digits))
	var f func(i int)
	f = func(i int) {
		if i == len(digits) {
			res = append(res, string(buf))
			return
		}
		for _, c := range []byte(keypad[digits[i]-'0']) {
			buf = append(buf, c)
			f(i + 1)
			buf = buf[:len(buf)-1]
		}
	}
	f(0)
	return res
}
