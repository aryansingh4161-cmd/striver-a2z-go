package xorrange

import "testing"

func TestXorRange(t *testing.T) {
	for l := 0; l <= 40; l++ {
		for r := l; r <= 40; r++ {
			want := 0
			for x := l; x <= r; x++ {
				want ^= x
			}
			if got := XorRange(l, r); got != want {
				t.Fatalf("XorRange(%d,%d) = %d want %d", l, r, got, want)
			}
		}
	}
}
