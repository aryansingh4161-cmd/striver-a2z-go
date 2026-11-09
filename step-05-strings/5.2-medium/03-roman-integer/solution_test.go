package romaninteger

import "testing"

func TestRoman(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{{"III", 3}, {"LVIII", 58}, {"MCMXCIV", 1994}, {"IV", 4}, {"MMMCMXCIX", 3999}}
	for _, tc := range tests {
		if got := RomanToInt(tc.s); got != tc.n {
			t.Errorf("RomanToInt(%q) = %d", tc.s, got)
		}
		if got := IntToRoman(tc.n); got != tc.s {
			t.Errorf("IntToRoman(%d) = %q", tc.n, got)
		}
	}
	for n := 1; n <= 3999; n++ { // round trip
		if RomanToInt(IntToRoman(n)) != n {
			t.Fatalf("round trip failed at %d", n)
		}
	}
}
