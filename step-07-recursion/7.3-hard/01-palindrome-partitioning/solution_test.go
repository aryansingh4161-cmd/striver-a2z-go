package palindromepartitioning

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	if got := Partition("aab"); !reflect.DeepEqual(got, [][]string{{"a", "a", "b"}, {"aa", "b"}}) {
		t.Errorf("got %v", got)
	}
	if got := Partition("a"); !reflect.DeepEqual(got, [][]string{{"a"}}) {
		t.Errorf("got %v", got)
	}
	if got := len(Partition("aaa")); got != 4 { // a a a | a aa | aa a | aaa
		t.Errorf("aaa gave %d", got)
	}
}
