package rabinkarp

const (
	base = 31
	mod  = 1_000_000_007
)

// Brute checks every start. O(n·m).
func Brute(s, p string) []int {
	out := []int{}
	for i := 0; i+len(p) <= len(s); i++ {
		if s[i:i+len(p)] == p {
			out = append(out, i)
		}
	}
	return out
}

func val(c byte) int { return int(c-'a') + 1 } // 1..26, so "a" and "aa" differ

// Search returns all match positions using a rolling hash. Expected O(n + m).
func Search(s, p string) []int {
	n, m := len(s), len(p)
	out := []int{}
	if m == 0 || m > n {
		return out
	}
	power := 1 // base^(m-1) mod mod
	for i := 1; i < m; i++ {
		power = power * base % mod
	}
	hp, hs := 0, 0
	for i := 0; i < m; i++ {
		hp = (hp*base + val(p[i])) % mod
		hs = (hs*base + val(s[i])) % mod
	}
	for i := 0; ; i++ {
		if hs == hp && s[i:i+m] == p { // verify to rule out collisions
			out = append(out, i)
		}
		if i+m == n {
			break
		}
		hs = (hs - val(s[i])*power%mod + mod) % mod // drop s[i]
		hs = (hs*base + val(s[i+m])) % mod          // add s[i+m]
	}
	return out
}
