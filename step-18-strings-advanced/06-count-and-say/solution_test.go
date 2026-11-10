package countandsay

import (
	"strings"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	want := []string{"1", "11", "21", "1211", "111221", "312211", "13112221", "1113213211"}
	for i, w := range want {
		if got := CountAndSay(i + 1); got != w {
			t.Errorf("n=%d got %q want %q", i+1, got, w)
		}
	}
	if s := CountAndSay(25); strings.ContainsAny(s, "0456789") {
		t.Error("only digits 1, 2, 3 should appear")
	}
}
