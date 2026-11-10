package binaryconversion

// ToBinary converts a non-negative int to its binary string. Time O(log n).
func ToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	var bits []byte
	for n > 0 {
		bits = append(bits, byte('0'+n%2))
		n /= 2
	}
	// remainders come out in reverse order
	for l, r := 0, len(bits)-1; l < r; l, r = l+1, r-1 {
		bits[l], bits[r] = bits[r], bits[l]
	}
	return string(bits)
}

// ToDecimal converts a binary string to an int. Time O(len).
func ToDecimal(s string) int {
	v := 0
	for i := 0; i < len(s); i++ {
		v = v*2 + int(s[i]-'0')
	}
	return v
}
