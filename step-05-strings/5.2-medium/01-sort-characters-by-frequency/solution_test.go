package sortcharactersbyfrequency

import "testing"

func TestFrequencySort(t *testing.T) {
	tests := []struct{ in, want string }{
		{"tree", "eert"}, {"cccaaa", "aaaccc"}, {"Aabb", "bbAa"}, {"", ""},
	}
	for _, tc := range tests {
		if got := FrequencySort(tc.in); got != tc.want {
			t.Errorf("FrequencySort(%q) = %q", tc.in, got)
		}
	}
}
