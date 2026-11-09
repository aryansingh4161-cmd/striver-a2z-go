package printnamentimes

import (
	"slices"
	"testing"
)

func TestPrintName(t *testing.T) {
	if got := PrintName("Aryan", 3); !slices.Equal(got, []string{"Aryan", "Aryan", "Aryan"}) {
		t.Errorf("got %v", got)
	}
	if got := PrintName("x", 0); len(got) != 0 {
		t.Errorf("n=0 should give nothing, got %v", got)
	}
}
