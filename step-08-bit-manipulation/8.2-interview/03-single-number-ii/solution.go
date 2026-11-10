package singlenumberii

// SingleBitCount counts each bit mod 3. Time O(64n), space O(1).
func SingleBitCount(a []int) int {
	res := 0
	for b := 0; b < 64; b++ {
		cnt := 0
		for _, x := range a {
			cnt += (x >> b) & 1
		}
		if cnt%3 != 0 {
			res |= 1 << b
		}
	}
	return res
}

// Single uses the ones/twos buckets. Time O(n), space O(1).
func Single(a []int) int {
	ones, twos := 0, 0
	for _, x := range a {
		ones = (ones ^ x) &^ twos
		twos = (twos ^ x) &^ ones
	}
	return ones
}
