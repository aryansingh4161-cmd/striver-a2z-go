package wordsearch

import "testing"

func grid() [][]byte {
	return [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
}

func TestExist(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{{"ABCCED", true}, {"SEE", true}, {"ABCB", false}, {"ASADFBCCEESE", true}, {"Z", false}}
	for _, tc := range tests {
		g := grid()
		if got := Exist(g, tc.word); got != tc.want {
			t.Errorf("Exist(%q) = %v", tc.word, got)
		}
		if string(g[0]) != "ABCE" || string(g[1]) != "SFCS" {
			t.Errorf("grid not restored after %q", tc.word)
		}
	}
}
