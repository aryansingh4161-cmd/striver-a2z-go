package validparentheses

// IsValid checks bracket matching with a stack. Time O(n), space O(n).
func IsValid(s string) bool {
	match := map[byte]byte{')': '(', ']': '[', '}': '{'}
	st := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			st = append(st, c)
		default:
			if len(st) == 0 || st[len(st)-1] != match[c] {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
