package countsetbits

import (
	"math/bits"
	"testing"
)

func TestCount(t *testing.T) {
	all := CountAll(2000)
	for n := 0; n <= 2000; n++ {
		want := bits.OnesCount(uint(n))
		if CountShift(n) != want || CountKernighan(n) != want || all[n] != want {
			t.Fatalf("n=%d", n)
		}
	}
}
