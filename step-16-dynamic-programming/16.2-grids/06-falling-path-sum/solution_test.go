package fallingpathsum

import (
	"math/rand"
	"testing"
)

func TestFalling(t *testing.T) {
	mat := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	for i, f := range []func([][]int, bool) int{Recursive, Tab, Optimal} {
		if got := f(mat, false); got != 13 {
			t.Errorf("min f%d = %d want 13", i, got)
		}
		if got := f(mat, true); got != 17 { // 3 → 5 → 9
			t.Errorf("max f%d = %d want 17", i, got)
		}
	}
	r := rand.New(rand.NewSource(10))
	for it := 0; it < 200; it++ {
		n, m := 1+r.Intn(5), 1+r.Intn(5)
		mat := make([][]int, n)
		for i := range mat {
			mat[i] = make([]int, m)
			for j := range mat[i] {
				mat[i][j] = r.Intn(41) - 20
			}
		}
		for _, wantMax := range []bool{false, true} {
			want := Recursive(mat, wantMax)
			if Tab(mat, wantMax) != want || Optimal(mat, wantMax) != want {
				t.Fatalf("%v max=%v", mat, wantMax)
			}
		}
	}
}
