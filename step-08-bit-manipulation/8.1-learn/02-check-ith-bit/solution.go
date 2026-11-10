package checkithbit

// IsSet reports whether bit i of n is 1. O(1).
func IsSet(n, i int) bool { return n&(1<<i) != 0 }

// IsSetShift is the same check using a right shift. O(1).
func IsSetShift(n, i int) bool { return (n>>i)&1 == 1 }

// Set returns n with bit i set to 1.
func Set(n, i int) int { return n | (1 << i) }

// Clear returns n with bit i set to 0.
func Clear(n, i int) int { return n &^ (1 << i) }

// Toggle returns n with bit i flipped.
func Toggle(n, i int) int { return n ^ (1 << i) }
