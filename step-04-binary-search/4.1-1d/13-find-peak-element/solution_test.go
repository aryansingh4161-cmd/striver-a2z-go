package findpeakelement

import (
	"math"
	"testing"
)

func isPeak(a []int, i int) bool {
	left, right := math.MinInt, math.MinInt
	if i > 0 {
		left = a[i-1]
	}
	if i < len(a)-1 {
		right = a[i+1]
	}
	return a[i] > left && a[i] > right
}

func TestPeak(t *testing.T) {
	cases := [][]int{{1, 2, 3, 1}, {1, 2, 1, 3, 5, 6, 4}, {1}, {2, 1}, {1, 2}, {5, 4, 3, 2, 1}}
	for _, a := range cases {
		if i := Peak(a); !isPeak(a, i) {
			t.Errorf("Peak(%v) = %d is not a peak", a, i)
		}
		if i := PeakLinear(a); !isPeak(a, i) {
			t.Errorf("PeakLinear(%v) = %d is not a peak", a, i)
		}
	}
}
