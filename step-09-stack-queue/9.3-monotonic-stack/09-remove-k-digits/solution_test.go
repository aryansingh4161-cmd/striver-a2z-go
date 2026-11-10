package removekdigits

import "testing"

func TestRemoveKDigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{{"1432219", 3, "1219"}, {"10200", 1, "200"}, {"10", 2, "0"}, {"12345", 2, "123"}, {"9", 1, "0"}, {"112", 1, "11"}}
	for _, tc := range tests {
		if got := RemoveKDigits(tc.num, tc.k); got != tc.want {
			t.Errorf("RemoveKDigits(%q,%d) = %q", tc.num, tc.k, got)
		}
	}
}
