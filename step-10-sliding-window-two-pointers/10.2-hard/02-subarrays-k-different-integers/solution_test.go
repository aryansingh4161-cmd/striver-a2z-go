package subarrayskdifferentintegers

import (
	"math/rand"
	"testing"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	if SubarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2) != 7 || SubarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3) != 3 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 300; it++ {
		a := make([]int, r.Intn(20))
		for i := range a {
			a[i] = r.Intn(4)
		}
		k := r.Intn(5)
		if SubarraysWithKDistinct(a, k) != SubarraysWithKDistinctBrute(a, k) {
			t.Fatalf("mismatch %v k=%d", a, k)
		}
	}
}
