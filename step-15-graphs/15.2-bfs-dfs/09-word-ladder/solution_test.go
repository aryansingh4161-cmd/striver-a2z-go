package wordladder

import "testing"

func TestLadderLength(t *testing.T) {
	list := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	if LadderLength("hit", "cog", list) != 5 {
		t.Error("example 1")
	}
	if LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}) != 0 {
		t.Error("end not in list")
	}
	if LadderLength("a", "c", []string{"a", "b", "c"}) != 2 {
		t.Error("one step")
	}
}
