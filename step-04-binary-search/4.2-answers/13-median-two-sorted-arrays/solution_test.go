package mediantwosortedarrays

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		a, b []int
		want float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 3, 4, 7, 10, 12}, []int{2, 3, 6, 15}, 5},
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
	}
	for _, tc := range tests {
		if got := Median(tc.a, tc.b); got != tc.want {
			t.Errorf("Median(%v,%v) = %v", tc.a, tc.b, got)
		}
		if got := MedianMerge(tc.a, tc.b); got != tc.want {
			t.Errorf("MedianMerge(%v,%v) = %v", tc.a, tc.b, got)
		}
	}
	r := rand.New(rand.NewSource(5))
	for k := 0; k < 200; k++ {
		a, b := randSorted(r), randSorted(r)
		if len(a)+len(b) == 0 {
			continue
		}
		if Median(a, b) != MedianMerge(a, b) {
			t.Fatalf("mismatch %v %v", a, b)
		}
	}
}

func randSorted(r *rand.Rand) []int {
	a := make([]int, r.Intn(8))
	for i := range a {
		a[i] = r.Intn(20) - 5
	}
	slices.Sort(a)
	return a
}
