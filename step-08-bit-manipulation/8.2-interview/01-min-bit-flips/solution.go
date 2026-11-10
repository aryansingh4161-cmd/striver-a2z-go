package minbitflips

// MinFlips counts differing bits between a and b. O(set bits).
func MinFlips(a, b int) int {
	x, c := a^b, 0
	for x != 0 {
		x &= x - 1
		c++
	}
	return c
}
