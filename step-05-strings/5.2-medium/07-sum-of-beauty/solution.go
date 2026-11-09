package sumofbeauty

// BeautySum returns the sum of (max freq - min freq) over all substrings.
// Time O(26 n^2), space O(1).
func BeautySum(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		var freq [26]int
		for j := i; j < len(s); j++ {
			freq[s[j]-'a']++
			hi, lo := 0, len(s)+1
			for _, f := range freq {
				if f > 0 {
					hi = max(hi, f)
					lo = min(lo, f)
				}
			}
			total += hi - lo
		}
	}
	return total
}
