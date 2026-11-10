package countgoodnumbers

const mod = 1_000_000_007

// CountGoodNumbers returns 5^ceil(n/2) * 4^floor(n/2) mod 1e9+7. Time O(log n).
func CountGoodNumbers(n int) int {
	return powMod(5, (n+1)/2) * powMod(4, n/2) % mod
}

// powMod computes base^exp % mod with binary exponentiation.
func powMod(base, exp int) int {
	base %= mod
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp /= 2
	}
	return result
}
