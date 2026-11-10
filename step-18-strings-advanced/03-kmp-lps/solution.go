// Package kmplps has the LPS array and KMP search, reused by Step 18.4 and 18.5.
package kmplps

// LPS returns the longest proper prefix-suffix lengths of p. O(m).
func LPS(p string) []int {
	lps := make([]int, len(p))
	length := 0
	for i := 1; i < len(p); i++ {
		for length > 0 && p[i] != p[length] {
			length = lps[length-1] // fall back, i stays
		}
		if p[i] == p[length] {
			length++
		}
		lps[i] = length
	}
	return lps
}

// Search returns every start index of p in s. O(n + m).
func Search(s, p string) []int {
	out := []int{}
	m := len(p)
	if m == 0 {
		return out
	}
	lps := LPS(p)
	j := 0
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != p[j] {
			j = lps[j-1]
		}
		if s[i] == p[j] {
			j++
		}
		if j == m {
			out = append(out, i-m+1)
			j = lps[j-1] // allow overlapping matches
		}
	}
	return out
}

// LPSBrute checks every prefix length directly. O(m³).
func LPSBrute(p string) []int {
	lps := make([]int, len(p))
	for i := range p {
		for l := i; l >= 1; l-- { // proper: l <= i
			if p[:l] == p[i-l+1:i+1] {
				lps[i] = l
				break
			}
		}
	}
	return lps
}
