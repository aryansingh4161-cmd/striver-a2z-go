package repeatingandmissing

import "testing"

func TestFind(t *testing.T) {
	tests := []struct {
		a    []int
		want [2]int
	}{
		{[]int{3, 1, 2, 5, 3}, [2]int{3, 4}},
		{[]int{1, 1}, [2]int{1, 2}},
		{[]int{2, 2}, [2]int{2, 1}},
		{[]int{4, 3, 6, 2, 1, 1}, [2]int{1, 5}},
	}
	for _, tc := range tests {
		for name, fn := range map[string]func([]int) [2]int{"hash": FindHash, "math": FindMath, "xor": FindXOR} {
			if got := fn(tc.a); got != tc.want {
				t.Errorf("%s(%v) = %v, want %v", name, tc.a, got, tc.want)
			}
		}
	}
}
