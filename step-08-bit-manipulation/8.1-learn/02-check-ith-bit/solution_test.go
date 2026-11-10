package checkithbit

import "testing"

func TestBits(t *testing.T) {
	n := 13 // 1101
	want := []bool{true, false, true, true, false}
	for i, w := range want {
		if IsSet(n, i) != w || IsSetShift(n, i) != w {
			t.Errorf("bit %d", i)
		}
	}
	if Set(13, 1) != 15 || Clear(13, 2) != 9 || Toggle(13, 0) != 12 || Toggle(13, 1) != 15 {
		t.Error("set/clear/toggle")
	}
}
