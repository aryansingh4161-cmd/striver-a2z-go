package mediandatastream

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	if m.FindMedian() != 1.5 {
		t.Error("1.5")
	}
	m.AddNum(3)
	if m.FindMedian() != 2 {
		t.Error("2")
	}
	r := rand.New(rand.NewSource(7))
	mf := Constructor()
	var all []int
	for i := 0; i < 500; i++ {
		x := r.Intn(100) - 50
		mf.AddNum(x)
		all = append(all, x)
		s := append([]int{}, all...)
		slices.Sort(s)
		var want float64
		if len(s)%2 == 1 {
			want = float64(s[len(s)/2])
		} else {
			want = float64(s[len(s)/2-1]+s[len(s)/2]) / 2
		}
		if mf.FindMedian() != want {
			t.Fatalf("step %d: %v want %v", i, mf.FindMedian(), want)
		}
	}
}
