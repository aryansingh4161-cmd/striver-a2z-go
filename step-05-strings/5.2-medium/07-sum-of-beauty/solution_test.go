package sumofbeauty

import "testing"

func TestBeautySum(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{{"aabcb", 5}, {"aabcbaa", 17}, {"a", 0}, {"abc", 0}}
	for _, tc := range tests {
		if got := BeautySum(tc.in); got != tc.want {
			t.Errorf("BeautySum(%q) = %d", tc.in, got)
		}
	}
}
