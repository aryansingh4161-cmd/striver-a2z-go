package removekdigits

// RemoveKDigits returns the smallest number after removing k digits. O(n).
func RemoveKDigits(num string, k int) string {
	st := []byte{}
	for i := 0; i < len(num); i++ {
		d := num[i]
		for k > 0 && len(st) > 0 && st[len(st)-1] > d {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, d)
	}
	st = st[:len(st)-min(k, len(st))] // still need to remove? take from the end
	i := 0
	for i < len(st) && st[i] == '0' { // strip leading zeros
		i++
	}
	if i == len(st) {
		return "0"
	}
	return string(st[i:])
}
