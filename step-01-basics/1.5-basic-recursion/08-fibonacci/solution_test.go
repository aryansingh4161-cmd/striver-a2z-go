package fibonacci

import "testing"

func TestFib(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for n, w := range want {
		if got := Fib(n); got != w {
			t.Errorf("Fib(%d) = %d, want %d", n, got, w)
		}
		if got := FibRecursive(n); got != w {
			t.Errorf("FibRecursive(%d) = %d, want %d", n, got, w)
		}
	}
	if got := Fib(50); got != 12586269025 {
		t.Errorf("Fib(50) = %d", got)
	}
}
