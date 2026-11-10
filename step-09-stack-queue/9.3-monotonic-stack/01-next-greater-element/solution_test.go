package nextgreaterelement

import (
	"math/rand"
	"slices"
	"testing"
)

func TestNGE(t *testing.T) {
	if got := NGE([]int{4, 5, 2, 10, 8}); !slices.Equal(got, []int{5, 10, 10, -1, -1}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(1))
	for k := 0; k < 100; k++ {
		a := make([]int, r.Intn(20))
		for i := range a {
			a[i] = r.Intn(10)
		}
		if !slices.Equal(NGE(a), NGEBrute(a)) {
			t.Fatalf("mismatch on %v", a)
		}
	}
}
