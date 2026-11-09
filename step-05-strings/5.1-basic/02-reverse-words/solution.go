package reversewords

import "strings"

// ReverseWordsFields uses strings.Fields. Time O(n).
func ReverseWordsFields(s string) string {
	words := strings.Fields(s)
	for l, r := 0, len(words)-1; l < r; l, r = l+1, r-1 {
		words[l], words[r] = words[r], words[l]
	}
	return strings.Join(words, " ")
}

// ReverseWords scans from the end without library splitting. Time O(n).
func ReverseWords(s string) string {
	var sb strings.Builder
	i := len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' { // skip spaces
			i--
		}
		if i < 0 {
			break
		}
		end := i
		for i >= 0 && s[i] != ' ' { // find the word start
			i--
		}
		if sb.Len() > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(s[i+1 : end+1])
	}
	return sb.String()
}
