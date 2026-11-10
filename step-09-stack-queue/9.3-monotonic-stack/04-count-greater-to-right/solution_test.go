package countgreatertoright

import (
	"math/rand"
	"slices"
	"testing"
)

func TestCount(t *testing.T) {
	a := []int{3, 4, 2, 7, 5, 8, 10, 6}
	if got := CountQueries(a, []int{0, 5}); !slices.Equal(got, []int{6, 1}) {
		t.Errorf("CountQueries = %v", got)
	}
	r := rand.New(rand.NewSource(2))
	for k := 0; k < 50; k++ {
		b := make([]int, r.Intn(25))
		all := make([]int, len(b))
		for i := range b {
			b[i] = r.Intn(10)
			all[i] = i
		}
		if !slices.Equal(CountAll(b), CountQueries(b, all)) {
			t.Fatalf("mismatch on %v", b)
		}
	}
}
