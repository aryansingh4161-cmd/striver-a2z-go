package sieve

import (
	"slices"
	"testing"
)

func TestSieve(t *testing.T) {
	if got := Primes(30); !slices.Equal(got, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}) {
		t.Errorf("got %v", got)
	}
	if CountPrimes(10) != 4 || CountPrimes(0) != 0 || CountPrimes(2) != 0 || CountPrimes(1_000_000) != 78498 {
		t.Error("CountPrimes")
	}
}
