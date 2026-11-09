package kthelementtwosortedarrays

import (
	"math/rand"
	"slices"
	"testing"
)

func TestKth(t *testing.T) {
	a, b := []int{2, 3, 6, 7, 9}, []int{1, 4, 8, 10}
	if got := Kth(a, b, 5); got != 6 {
		t.Errorf("Kth = %d", got)
	}
	r := rand.New(rand.NewSource(9))
	for it := 0; it < 300; it++ {
		x, y := randSorted(r), randSorted(r)
		all := append(append([]int{}, x...), y...)
		slices.Sort(all)
		for k := 1; k <= len(all); k++ {
			if got := Kth(x, y, k); got != all[k-1] {
				t.Fatalf("Kth(%v,%v,%d) = %d want %d", x, y, k, got, all[k-1])
			}
			if got := KthMerge(x, y, k); got != all[k-1] {
				t.Fatalf("KthMerge(%v,%v,%d) = %d", x, y, k, got)
			}
		}
	}
}

func randSorted(r *rand.Rand) []int {
	a := make([]int, r.Intn(7))
	for i := range a {
		a[i] = r.Intn(15)
	}
	slices.Sort(a)
	return a
}
