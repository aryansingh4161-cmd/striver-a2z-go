package generateparentheses

import (
	"slices"
	"testing"
)

func TestGenerate(t *testing.T) {
	if got := Generate(3); !slices.Equal(got, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}) {
		t.Errorf("Generate(3) = %v", got)
	}
	if got := Generate(1); !slices.Equal(got, []string{"()"}) {
		t.Errorf("Generate(1) = %v", got)
	}
	catalan := []int{1, 1, 2, 5, 14, 42}
	for n := 1; n <= 5; n++ {
		if got := len(Generate(n)); got != catalan[n] {
			t.Errorf("len(Generate(%d)) = %d", n, got)
		}
	}
}
