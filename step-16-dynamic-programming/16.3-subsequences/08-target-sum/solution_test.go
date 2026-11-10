package targetsum

import (
	"math/rand"
	"testing"
)

func TestTargetSum(t *testing.T) {
	if Ways([]int{1, 1, 1, 1, 1}, 3) != 5 || Recursive([]int{1, 1, 1, 1, 1}, 3) != 5 {
		t.Error("example")
	}
	if Ways([]int{1}, 2) != 0 {
		t.Error("impossible")
	}
	if Ways([]int{0, 0, 1}, 1) != 4 {
		t.Error("zeros double the count")
	}
	r := rand.New(rand.NewSource(8))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = r.Intn(5)
		}
		target := r.Intn(21) - 10
		if got, want := Ways(a, target), Recursive(a, target); got != want {
			t.Fatalf("%v target %d: %d vs %d", a, target, got, want)
		}
	}
}
