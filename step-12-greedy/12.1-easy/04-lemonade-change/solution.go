package lemonadechange

// LemonadeChange reports whether everyone can get correct change. O(n), O(1).
func LemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		default: // 20
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
