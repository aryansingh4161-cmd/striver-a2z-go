package patterns

import "testing"

func TestPatterns(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"1", Pattern1(3), "***\n***\n***"},
		{"2", Pattern2(3), "*\n**\n***"},
		{"3", Pattern3(3), "1\n12\n123"},
		{"4", Pattern4(3), "1\n22\n333"},
		{"5", Pattern5(3), "***\n**\n*"},
		{"6", Pattern6(3), "123\n12\n1"},
		{"7", Pattern7(3), "  *\n ***\n*****"},
		{"8", Pattern8(3), "*****\n ***\n  *"},
		{"9", Pattern9(2), " *\n***\n***\n *"},
		{"10", Pattern10(3), "*\n**\n***\n**\n*"},
		{"11", Pattern11(3), "1\n01\n101"},
		{"12", Pattern12(3), "1    1\n12  21\n123321"},
		{"13", Pattern13(3), "1\n2 3\n4 5 6"},
		{"14", Pattern14(3), "A\nAB\nABC"},
		{"15", Pattern15(3), "ABC\nAB\nA"},
		{"16", Pattern16(3), "A\nBB\nCCC"},
		{"17", Pattern17(3), "  A\n ABA\nABCBA"},
		{"18", Pattern18(3), "C\nBC\nABC"},
		{"19", Pattern19(3), "******\n**  **\n*    *\n*    *\n**  **\n******"},
		{"20", Pattern20(3), "*    *\n**  **\n******\n**  **\n*    *"},
		{"21", Pattern21(4), "****\n*  *\n*  *\n****"},
		{"22", Pattern22(3), "33333\n32223\n32123\n32223\n33333"},
	}
	for _, tc := range tests {
		if tc.got != tc.want {
			t.Errorf("Pattern%s:\ngot:\n%s\nwant:\n%s", tc.name, tc.got, tc.want)
		}
	}
}
