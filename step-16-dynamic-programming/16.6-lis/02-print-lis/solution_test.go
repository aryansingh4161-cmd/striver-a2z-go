package printlis

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.6-lis/01-lis"
)

func isSubseq(sub, a []int) bool {
	i := 0
	for _, x := range a {
		if i < len(sub) && sub[i] == x {
			i++
		}
	}
	return i == len(sub)
}

func TestPrintLIS(t *testing.T) {
	if got := PrintLIS([]int{5, 4, 11, 1, 16, 8}); !slices.Equal(got, []int{5, 11, 16}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(12))
		for i := range a {
			a[i] = r.Intn(10)
		}
		got := PrintLIS(a)
		if len(got) != lis.EndingAt(a) || !isSubseq(got, a) {
			t.Fatalf("%v → %v", a, got)
		}
		for i := 1; i < len(got); i++ {
			if got[i] <= got[i-1] {
				t.Fatalf("not increasing: %v", got)
			}
		}
	}
}
