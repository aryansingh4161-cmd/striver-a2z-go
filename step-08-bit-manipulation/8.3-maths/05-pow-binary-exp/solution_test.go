package powbinaryexp

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	if math.Abs(Pow(2, 10)-1024) > 1e-9 || math.Abs(Pow(2, -2)-0.25) > 1e-9 || Pow(3, 0) != 1 {
		t.Error("Pow")
	}
	const m = 1_000_000_007
	if PowMod(2, 10, m) != 1024 || PowMod(2, 1_000_000_006, m) != 1 { // Fermat: 2^(p-1) = 1 mod p
		t.Error("PowMod")
	}
	if PowMod(5, 0, 7) != 1 || PowMod(10, 3, 1) != 0 {
		t.Error("PowMod edge cases")
	}
}
