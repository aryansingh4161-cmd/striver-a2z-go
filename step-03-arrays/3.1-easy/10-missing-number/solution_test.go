package missingnumber

import "testing"

func TestMissing(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{3, 0, 1}, 2}, {[]int{0, 1}, 2}, {[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8}, {[]int{1}, 0}, {[]int{0}, 1}}
	for _, tc := range tests {
		for name, fn := range map[string]func([]int) int{"hash": MissingHash, "sum": MissingSum, "xor": MissingXOR} {
			if got := fn(tc.a); got != tc.want {
				t.Errorf("%s(%v) = %d, want %d", name, tc.a, got, tc.want)
			}
		}
	}
}
