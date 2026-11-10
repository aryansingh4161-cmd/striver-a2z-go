package slidingwindowmaximum

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	if got := MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !slices.Equal(got, []int{3, 3, 5, 5, 6, 7}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 200; it++ {
		a := make([]int, r.Intn(20)+1)
		for i := range a {
			a[i] = r.Intn(10)
		}
		k := r.Intn(len(a)) + 1
		if !slices.Equal(MaxSlidingWindow(a, k), MaxBrute(a, k)) {
			t.Fatalf("mismatch %v k=%d", a, k)
		}
	}
}
