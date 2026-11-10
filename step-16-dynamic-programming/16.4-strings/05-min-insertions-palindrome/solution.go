package mininsertionspalindrome

import lps "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.4-strings/04-longest-palindromic-subsequence"

// MinInsertions = n − LPS(s). O(n²).
func MinInsertions(s string) int { return len(s) - lps.LPS(s) }
