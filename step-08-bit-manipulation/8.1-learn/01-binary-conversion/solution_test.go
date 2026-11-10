package binaryconversion

import (
	"strconv"
	"testing"
)

func TestConversion(t *testing.T) {
	for n := 0; n <= 1000; n++ {
		b := ToBinary(n)
		if b != strconv.FormatInt(int64(n), 2) {
			t.Fatalf("ToBinary(%d) = %s", n, b)
		}
		if ToDecimal(b) != n {
			t.Fatalf("ToDecimal(%s) = %d", b, ToDecimal(b))
		}
	}
}
