package rotatestring

import "strings"

// RotateBrute tries every rotation. Time O(n^2).
func RotateBrute(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	for k := 1; k < len(s); k++ {
		if s[k:]+s[:k] == goal {
			return true
		}
	}
	return false
}

// Rotate checks whether goal appears in s+s.
func Rotate(s, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
