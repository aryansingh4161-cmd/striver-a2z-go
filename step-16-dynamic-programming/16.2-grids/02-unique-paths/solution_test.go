package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	fs := map[string]func(int, int) int{"rec": Recursive, "memo": Memo, "tab": Tab, "opt": Optimal, "comb": Combinatorics}
	for m := 1; m <= 8; m++ {
		for n := 1; n <= 8; n++ {
			want := Recursive(m, n)
			for name, f := range fs {
				if got := f(m, n); got != want {
					t.Errorf("%s(%d,%d) = %d want %d", name, m, n, got, want)
				}
			}
		}
	}
	if Combinatorics(3, 7) != 28 || Optimal(23, 12) != 193536720 {
		t.Error("known values")
	}
}
