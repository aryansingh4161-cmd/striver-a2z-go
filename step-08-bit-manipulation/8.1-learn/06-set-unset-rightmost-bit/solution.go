package rightmostbit

// ClearLowestSet turns off the rightmost 1 bit.
func ClearLowestSet(n int) int { return n & (n - 1) }

// SetLowestUnset turns on the rightmost 0 bit.
func SetLowestUnset(n int) int { return n | (n + 1) }

// LowestSetBit keeps only the rightmost 1 bit.
func LowestSetBit(n int) int { return n & -n }
