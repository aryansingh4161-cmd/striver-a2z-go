package pow

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{{2, 10, 1024}, {2.1, 3, 9.261}, {2, -2, 0.25}, {5, 0, 1}, {1, -2147483648, 1}, {-2, 3, -8}}
	for _, tc := range tests {
		for name, fn := range map[string]func(float64, int) float64{"rec": PowRecursive, "iter": Pow} {
			if got := fn(tc.x, tc.n); math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("%s(%v,%d) = %v", name, tc.x, tc.n, got)
			}
		}
	}
}
