package maxsumcombinations

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

func TestTopKSums(t *testing.T) {
	if got := TopKSums([]int{3, 2}, []int{1, 4}, 2); !slices.Equal(got, []int{7, 6}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 100; it++ {
		n := r.Intn(6) + 1
		a, b := make([]int, n), make([]int, n)
		var all []int
		for i := range a {
			a[i], b[i] = r.Intn(10), r.Intn(10)
		}
		for _, x := range a {
			for _, y := range b {
				all = append(all, x+y)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(all)))
		k := r.Intn(n*n) + 1
		if got := TopKSums(a, b, k); !slices.Equal(got, all[:k]) {
			t.Fatalf("a=%v b=%v k=%d got %v want %v", a, b, k, got, all[:k])
		}
	}
}
