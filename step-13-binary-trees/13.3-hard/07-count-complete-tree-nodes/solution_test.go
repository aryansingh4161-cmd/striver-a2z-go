package countcompletetreenodes

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestCountNodes(t *testing.T) {
	for n := 0; n <= 70; n++ {
		vals := make([]int, n)
		for i := range vals {
			vals[i] = i + 1
		}
		if got := CountNodes(tree.FromLevel(vals)); got != n {
			t.Fatalf("n=%d got %d", n, got)
		}
	}
}
