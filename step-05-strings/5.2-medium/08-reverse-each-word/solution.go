package reverseeachword

// ReverseEachWord reverses the letters of each word, keeping order and spaces.
// Time O(n), space O(n).
func ReverseEachWord(s string) string {
	b := []byte(s)
	i := 0
	for i < len(b) {
		for i < len(b) && b[i] == ' ' {
			i++
		}
		j := i
		for j < len(b) && b[j] != ' ' {
			j++
		}
		for l, r := i, j-1; l < r; l, r = l+1, r-1 {
			b[l], b[r] = b[r], b[l]
		}
		i = j
	}
	return string(b)
}
