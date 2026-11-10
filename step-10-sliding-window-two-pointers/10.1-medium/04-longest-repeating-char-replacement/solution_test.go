package longestrepeatingcharreplacement

import (
	"math/rand"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	if CharacterReplacement("ABAB", 2) != 4 || CharacterReplacement("AABABBA", 1) != 4 || CharacterReplacement("", 3) != 0 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(20))
		for i := range b {
			b[i] = byte('A' + r.Intn(3))
		}
		k := r.Intn(4)
		if CharacterReplacement(string(b), k) != CharacterReplacementBrute(string(b), k) {
			t.Fatalf("mismatch %q k=%d", b, k)
		}
	}
}
