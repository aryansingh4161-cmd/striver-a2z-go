package trappingrainwater

import (
	"math/rand"
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		h    []int
		want int
	}{{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}, {[]int{4, 2, 0, 3, 2, 5}, 9}, {[]int{}, 0}, {[]int{3}, 0}, {[]int{1, 2, 3}, 0}}
	for _, tc := range tests {
		for name, fn := range map[string]func([]int) int{"prefix": TrapPrefix, "two": Trap, "stack": TrapStack} {
			if got := fn(tc.h); got != tc.want {
				t.Errorf("%s(%v) = %d", name, tc.h, got)
			}
		}
	}
	r := rand.New(rand.NewSource(3))
	for k := 0; k < 100; k++ {
		h := make([]int, r.Intn(30))
		for i := range h {
			h[i] = r.Intn(8)
		}
		a, b, c := TrapPrefix(h), Trap(h), TrapStack(h)
		if a != b || b != c {
			t.Fatalf("mismatch on %v: %d %d %d", h, a, b, c)
		}
	}
}
