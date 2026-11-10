package mininsertdelete

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.4-strings/01-lcs"

// MinOperations = n + m − 2·LCS. O(n·m).
func MinOperations(a, b string) int {
	return len(a) + len(b) - 2*lcs.Optimal(a, b)
}
