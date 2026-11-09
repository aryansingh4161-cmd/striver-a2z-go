package romaninteger

import "strings"

var value = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

// RomanToInt converts a Roman numeral to an integer. Time O(n).
func RomanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := value[s[i]]
		if i+1 < len(s) && v < value[s[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	return total
}

// IntToRoman converts 1..3999 to a Roman numeral using a greedy table.
func IntToRoman(n int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb strings.Builder
	for i, v := range vals {
		for n >= v {
			sb.WriteString(syms[i])
			n -= v
		}
	}
	return sb.String()
}
