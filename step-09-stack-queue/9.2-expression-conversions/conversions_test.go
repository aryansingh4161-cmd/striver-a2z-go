package conversions

import "testing"

func TestConversions(t *testing.T) {
	tests := []struct{ infix, postfix, prefix string }{
		{"(a+b)*c", "ab+c*", "*+abc"},
		{"a+b*(c^d-e)", "abcd^e-*+", "+a*b-^cde"},
		{"a-b-c", "ab-c-", "--abc"},
		{"a^b^c", "abc^^", "^a^bc"},
		{"a*(b+c)/d", "abc+*d/", "/*a+bcd"},
	}
	for _, tc := range tests {
		if got := InfixToPostfix(tc.infix); got != tc.postfix {
			t.Errorf("InfixToPostfix(%q) = %q", tc.infix, got)
		}
		if got := InfixToPrefix(tc.infix); got != tc.prefix {
			t.Errorf("InfixToPrefix(%q) = %q", tc.infix, got)
		}
		if got := PostfixToPrefix(tc.postfix); got != tc.prefix {
			t.Errorf("PostfixToPrefix(%q) = %q", tc.postfix, got)
		}
		if got := PrefixToPostfix(tc.prefix); got != tc.postfix {
			t.Errorf("PrefixToPostfix(%q) = %q", tc.prefix, got)
		}
		// infix from both forms must agree, and convert back to the same postfix
		pi, qi := PostfixToInfix(tc.postfix), PrefixToInfix(tc.prefix)
		if pi != qi {
			t.Errorf("PostfixToInfix %q vs PrefixToInfix %q", pi, qi)
		}
		if got := InfixToPostfix(pi); got != tc.postfix {
			t.Errorf("round trip %q -> %q", pi, got)
		}
	}
	if got := PostfixToInfix("ab+c*"); got != "((a+b)*c)" {
		t.Errorf("PostfixToInfix = %q", got)
	}
}
