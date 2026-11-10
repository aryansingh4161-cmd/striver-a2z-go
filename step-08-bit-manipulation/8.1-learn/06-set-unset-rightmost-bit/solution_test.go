package rightmostbit

import "testing"

func TestRightmost(t *testing.T) {
	if ClearLowestSet(12) != 8 || ClearLowestSet(7) != 6 || ClearLowestSet(0) != 0 {
		t.Error("ClearLowestSet")
	}
	if SetLowestUnset(11) != 15 || SetLowestUnset(10) != 11 || SetLowestUnset(0) != 1 {
		t.Error("SetLowestUnset")
	}
	if LowestSetBit(12) != 4 || LowestSetBit(7) != 1 || LowestSetBit(8) != 8 {
		t.Error("LowestSetBit")
	}
}
