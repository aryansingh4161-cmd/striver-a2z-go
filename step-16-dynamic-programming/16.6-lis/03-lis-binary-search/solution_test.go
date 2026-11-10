package lisbinarysearch

import (
	"math/rand"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.6-lis/01-lis"
)

func TestLIS(t *testing.T) {
	if LIS([]int{1, 7, 8, 4, 5, 6, -1, 9}) != 5 || LIS([]int{7, 7, 7}) != 1 || LIS(nil) != 0 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 500; it++ {
		a := make([]int, 1+r.Intn(30))
		for i := range a {
			a[i] = r.Intn(15)
		}
		if got, want := LIS(a), lis.EndingAt(a); got != want {
			t.Fatalf("%v: %d vs %d", a, got, want)
		}
	}
}
