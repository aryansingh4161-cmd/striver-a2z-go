package binarysubarrayswithsum

import (
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {
	if Count([]int{1, 0, 1, 0, 1}, 2) != 4 || Count([]int{0, 0, 0, 0, 0}, 0) != 15 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 300; it++ {
		a := make([]int, r.Intn(20))
		for i := range a {
			a[i] = r.Intn(2)
		}
		g := r.Intn(4)
		if Count(a, g) != CountPrefix(a, g) {
			t.Fatalf("mismatch %v goal=%d", a, g)
		}
	}
}
