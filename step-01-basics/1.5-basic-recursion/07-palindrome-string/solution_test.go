package palindromestring

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{{"madam", true}, {"hello", false}, {"", true}, {"a", true}, {"abba", true}, {"ab", false}}
	for _, tc := range tests {
		if got := IsPalindrome(tc.s); got != tc.want {
			t.Errorf("IsPalindrome(%q) = %v", tc.s, got)
		}
	}
	if !IsPalindromeClean("A man, a plan, a canal: Panama") {
		t.Error("Panama should be a palindrome")
	}
	if IsPalindromeClean("race a car") {
		t.Error("race a car should not be a palindrome")
	}
}
