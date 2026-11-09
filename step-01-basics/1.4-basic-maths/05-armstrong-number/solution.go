package armstrongnumber

// IsArmstrong reports whether n equals the sum of its digits each raised
// to the number of digits.
func IsArmstrong(n int) bool {
	if n < 0 {
		return false
	}
	// step 1: count digits
	k, tmp := 0, n
	for {
		k++
		tmp /= 10
		if tmp == 0 {
			break
		}
	}
	// step 2: sum digit^k
	sum, tmp := 0, n
	for tmp > 0 {
		sum += power(tmp%10, k)
		tmp /= 10
	}
	return sum == n
}

// power computes base^exp with integers (math.Pow uses float64).
func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
