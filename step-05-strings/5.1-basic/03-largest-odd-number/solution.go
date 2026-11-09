package largestoddnumber

// LargestOdd returns the largest odd number that is a prefix-substring of num.
// Time O(n), space O(1).
func LargestOdd(num string) string {
	start := 0
	for start < len(num)-1 && num[start] == '0' { // drop leading zeros
		start++
	}
	for i := len(num) - 1; i >= start; i-- {
		if (num[i]-'0')%2 == 1 {
			return num[start : i+1]
		}
	}
	return ""
}
