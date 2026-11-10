package poweroftwo

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	for n := -4; n <= 1100; n++ {
		want := false
		for p := 1; p <= n; p *= 2 {
			if p == n {
				want = true
			}
		}
		if IsPowerOfTwo(n) != want || IsPowerOfTwoLoop(n) != want {
			t.Fatalf("n=%d", n)
		}
	}
}
