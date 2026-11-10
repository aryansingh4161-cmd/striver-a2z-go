package sumsubarrayminimums

import (
	"math/rand"
	"testing"
)

func TestSumMins(t *testing.T) {
	if SumMins([]int{3, 1, 2, 4}) != 17 || SumMins([]int{11, 81, 94, 43, 3}) != 444 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(4))
	for k := 0; k < 100; k++ {
		a := make([]int, r.Intn(25)+1)
		for i := range a {
			a[i] = r.Intn(6) + 1 // small range → many ties
		}
		if SumMins(a) != SumMinsBrute(a) {
			t.Fatalf("mismatch on %v", a)
		}
	}
}
