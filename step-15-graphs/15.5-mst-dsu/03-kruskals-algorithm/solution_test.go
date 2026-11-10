package kruskalsalgorithm

import (
	"math/rand"
	"testing"

	primsalgorithm "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/01-prims-algorithm"
)

func TestMST(t *testing.T) {
	edges := [][3]int{{0, 1, 2}, {0, 3, 6}, {1, 2, 3}, {1, 3, 8}, {1, 4, 5}, {2, 4, 7}}
	if got := MST(5, edges); got != 16 {
		t.Errorf("got %d", got)
	}
}

// Prim and Kruskal must agree on random connected graphs.
func TestAgainstPrim(t *testing.T) {
	r := rand.New(rand.NewSource(14))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(9)
		var edges [][3]int
		for v := 1; v < n; v++ { // random tree keeps it connected
			edges = append(edges, [3]int{r.Intn(v), v, 1 + r.Intn(20)})
		}
		for k := r.Intn(10); k > 0 && n > 1; k-- {
			edges = append(edges, [3]int{r.Intn(n), r.Intn(n), 1 + r.Intn(20)})
		}
		want, _ := primsalgorithm.MST(n, edges)
		if got := MST(n, edges); got != want {
			t.Fatalf("kruskal %d prim %d", got, want)
		}
	}
}
