package wordladderii

import (
	"reflect"
	"testing"
)

func TestFindLadders(t *testing.T) {
	got := FindLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	want := [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
	if got := FindLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}); len(got) != 0 {
		t.Errorf("got %v", got)
	}
	got = FindLadders("a", "c", []string{"a", "b", "c"})
	if !reflect.DeepEqual(got, [][]string{{"a", "c"}}) {
		t.Errorf("got %v", got)
	}
}
