package subsequencesumkexists

import "testing"

func TestFind(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want bool
	}{{[]int{1, 2, 3}, 5, true}, {[]int{1, 2, 3}, 7, false}, {[]int{4}, 0, true}, {[]int{-1, 2}, 1, true}}
	for _, tc := range tests {
		path, ok := Find(tc.a, tc.k)
		if ok != tc.want {
			t.Errorf("Find(%v,%d) ok = %v", tc.a, tc.k, ok)
			continue
		}
		if ok {
			s := 0
			for _, x := range path {
				s += x
			}
			if s != tc.k {
				t.Errorf("Find(%v,%d) path %v sums to %d", tc.a, tc.k, path, s)
			}
		}
	}
}
