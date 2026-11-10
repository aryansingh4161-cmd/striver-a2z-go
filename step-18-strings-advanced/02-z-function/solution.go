package zfunction

// ZBrute extends from every position. O(n²).
func ZBrute(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i := 1; i < n; i++ {
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
	}
	return z
}

// Z builds the Z-array with the Z-box trick. O(n).
func Z(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0 // s[l:r] matches a prefix of s
	for i := 1; i < n; i++ {
		if i < r {
			z[i] = min(r-i, z[i-l]) // reuse what we know inside the box
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i] > r {
			l, r = i, i+z[i]
		}
	}
	return z
}

// Search returns all positions where p occurs in s, using Z on p + "$" + s.
// p and s must not contain '$'. O(n + m).
func Search(s, p string) []int {
	m := len(p)
	out := []int{}
	if m == 0 {
		return out
	}
	z := Z(p + "$" + s)
	for i := m + 1; i < len(z); i++ {
		if z[i] == m {
			out = append(out, i-m-1)
		}
	}
	return out
}
