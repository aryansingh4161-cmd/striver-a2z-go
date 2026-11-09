package palindromenumber

// IsPalindrome reports whether n reads the same backwards. Time O(log n), space O(1).
func IsPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	original := n // save a copy: the loop below turns n into 0
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev == original
}
