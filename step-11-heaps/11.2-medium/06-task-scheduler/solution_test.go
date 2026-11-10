package taskscheduler

import (
	"math/rand"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		tasks   string
		n, want int
	}{{"AAABBB", 2, 8}, {"ACABDB", 1, 6}, {"AAABBB", 3, 10}, {"AAABBB", 0, 6}, {"AAAAAABCDEFG", 2, 16}}
	for _, tc := range tests {
		if got := LeastIntervalFormula([]byte(tc.tasks), tc.n); got != tc.want {
			t.Errorf("Formula(%s,%d) = %d", tc.tasks, tc.n, got)
		}
		if got := LeastIntervalHeap([]byte(tc.tasks), tc.n); got != tc.want {
			t.Errorf("Heap(%s,%d) = %d", tc.tasks, tc.n, got)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(20)+1)
		for i := range b {
			b[i] = byte('A' + r.Intn(5))
		}
		n := r.Intn(4)
		if LeastIntervalFormula(b, n) != LeastIntervalHeap(b, n) {
			t.Fatalf("mismatch %s n=%d", b, n)
		}
	}
}
