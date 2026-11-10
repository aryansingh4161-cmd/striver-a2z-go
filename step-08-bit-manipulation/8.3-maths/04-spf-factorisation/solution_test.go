package spffactorisation

import (
	"slices"
	"testing"
)

func TestFactorise(t *testing.T) {
	spf := NewSPF(1000)
	tests := []struct {
		x    int
		want []int
	}{{12, []int{2, 2, 3}}, {97, []int{97}}, {1, []int{}}, {360, []int{2, 2, 2, 3, 3, 5}}, {999, []int{3, 3, 3, 37}}}
	for _, tc := range tests {
		if got := spf.Factorise(tc.x); !slices.Equal(got, tc.want) {
			t.Errorf("Factorise(%d) = %v", tc.x, got)
		}
	}
	for x := 2; x <= 1000; x++ { // product of factors must give x back
		p := 1
		for _, f := range spf.Factorise(x) {
			p *= f
		}
		if p != x {
			t.Fatalf("x=%d", x)
		}
	}
}
