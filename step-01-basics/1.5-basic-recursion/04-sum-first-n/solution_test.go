package sumfirstn

import "testing"

func TestSum(t *testing.T) {
	for _, tc := range []struct{ n, want int }{{3, 6}, {5, 15}, {0, 0}, {1, 1}, {100, 5050}} {
		for name, fn := range map[string]func(int) int{"param": SumParam, "func": SumFunc, "formula": SumFormula} {
			if got := fn(tc.n); got != tc.want {
				t.Errorf("%s(%d) = %d, want %d", name, tc.n, got, tc.want)
			}
		}
	}
}
