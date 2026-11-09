package countsubstringskdistinct

// CountBrute checks all substrings. Time O(n^2).
func CountBrute(s string, k int) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		var cnt [26]int
		distinct := 0
		for j := i; j < len(s); j++ {
			if cnt[s[j]-'a'] == 0 {
				distinct++
			}
			cnt[s[j]-'a']++
			if distinct == k {
				ans++
			} else if distinct > k {
				break
			}
		}
	}
	return ans
}

// Count returns atMost(k) - atMost(k-1). Time O(n), space O(1).
func Count(s string, k int) int {
	return atMost(s, k) - atMost(s, k-1)
}

func atMost(s string, k int) int {
	if k < 0 {
		return 0
	}
	var cnt [26]int
	l, distinct, total := 0, 0, 0
	for r := 0; r < len(s); r++ {
		if cnt[s[r]-'a'] == 0 {
			distinct++
		}
		cnt[s[r]-'a']++
		for distinct > k {
			cnt[s[l]-'a']--
			if cnt[s[l]-'a'] == 0 {
				distinct--
			}
			l++
		}
		total += r - l + 1
	}
	return total
}
