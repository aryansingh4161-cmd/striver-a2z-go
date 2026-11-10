// Package conversions converts between infix, prefix and postfix expressions.
// Operands are single letters/digits; operators are + - * / ^.
package conversions

import "strings"

func isOperand(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func prec(c byte) int {
	switch c {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return -1
}

// InfixToPostfix uses the Shunting-yard algorithm. O(n).
func InfixToPostfix(s string) string {
	return toPostfix(s, false)
}

// toPostfix is shunting-yard. reversed=true is used by InfixToPrefix, where
// equal-precedence operators are only popped for '^'.
func toPostfix(s string, reversed bool) string {
	var out strings.Builder
	st := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case isOperand(c):
			out.WriteByte(c)
		case c == '(':
			st = append(st, c)
		case c == ')':
			for len(st) > 0 && st[len(st)-1] != '(' {
				out.WriteByte(st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = st[:len(st)-1] // drop '('
		default: // operator
			for len(st) > 0 && st[len(st)-1] != '(' {
				top := st[len(st)-1]
				popEqual := c != '^' // left-associative
				if reversed {
					popEqual = c == '^'
				}
				if prec(top) > prec(c) || (prec(top) == prec(c) && popEqual) {
					out.WriteByte(top)
					st = st[:len(st)-1]
				} else {
					break
				}
			}
			st = append(st, c)
		}
	}
	for len(st) > 0 {
		out.WriteByte(st[len(st)-1])
		st = st[:len(st)-1]
	}
	return out.String()
}

// InfixToPrefix: reverse (swapping brackets), shunting-yard, reverse. O(n).
func InfixToPrefix(s string) string {
	b := []byte(s)
	reverse(b)
	for i, c := range b {
		if c == '(' {
			b[i] = ')'
		} else if c == ')' {
			b[i] = '('
		}
	}
	p := []byte(toPostfix(string(b), true))
	reverse(p)
	return string(p)
}

// PostfixToInfix builds fully bracketed infix. O(n) stack operations.
func PostfixToInfix(s string) string {
	st := []string{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isOperand(c) {
			st = append(st, string(c))
			continue
		}
		b, a := st[len(st)-1], st[len(st)-2]
		st = st[:len(st)-2]
		st = append(st, "("+a+string(c)+b+")")
	}
	return st[0]
}

// PrefixToInfix scans right to left. O(n) stack operations.
func PrefixToInfix(s string) string {
	st := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if isOperand(c) {
			st = append(st, string(c))
			continue
		}
		a, b := st[len(st)-1], st[len(st)-2]
		st = st[:len(st)-2]
		st = append(st, "("+a+string(c)+b+")")
	}
	return st[0]
}

// PostfixToPrefix: operator → op + a + b. O(n) stack operations.
func PostfixToPrefix(s string) string {
	st := []string{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isOperand(c) {
			st = append(st, string(c))
			continue
		}
		b, a := st[len(st)-1], st[len(st)-2]
		st = st[:len(st)-2]
		st = append(st, string(c)+a+b)
	}
	return st[0]
}

// PrefixToPostfix: right to left, operator → a + b + op. O(n) stack operations.
func PrefixToPostfix(s string) string {
	st := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if isOperand(c) {
			st = append(st, string(c))
			continue
		}
		a, b := st[len(st)-1], st[len(st)-2]
		st = st[:len(st)-2]
		st = append(st, a+b+string(c))
	}
	return st[0]
}

func reverse(b []byte) {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
}
