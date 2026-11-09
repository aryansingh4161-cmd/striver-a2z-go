// Package patterns solves Striver's 22 pattern problems.
//
// Every function returns the pattern as a string (rows joined by "\n")
// instead of printing it. That makes it easy to test. To see a pattern,
// call fmt.Println(patterns.Pattern7(5)).
//
// Trailing spaces at the end of each row are removed.
package patterns

import "strings"

// builder helps collect rows and join them at the end.
type builder struct{ rows []string }

func (b *builder) add(row string) { b.rows = append(b.rows, strings.TrimRight(row, " ")) }
func (b *builder) String() string { return strings.Join(b.rows, "\n") }

// rep repeats s n times (n may be 0).
func rep(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}

// Pattern1: n x n square of stars.
//
//	***
//	***
//	***
func Pattern1(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		b.add(rep("*", n))
	}
	return b.String()
}

// Pattern2: right triangle of stars. Row i has i+1 stars.
//
//	*
//	**
//	***
func Pattern2(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		b.add(rep("*", i+1))
	}
	return b.String()
}

// Pattern3: right triangle of numbers 1..row.
//
//	1
//	12
//	123
func Pattern3(n int) string {
	var b builder
	for i := 1; i <= n; i++ {
		row := ""
		for j := 1; j <= i; j++ {
			row += string(rune('0' + j%10))
		}
		b.add(row)
	}
	return b.String()
}

// Pattern4: row number repeated.
//
//	1
//	22
//	333
func Pattern4(n int) string {
	var b builder
	for i := 1; i <= n; i++ {
		b.add(rep(string(rune('0'+i%10)), i))
	}
	return b.String()
}

// Pattern5: inverted right triangle of stars.
//
//	***
//	**
//	*
func Pattern5(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		b.add(rep("*", n-i))
	}
	return b.String()
}

// Pattern6: inverted number triangle.
//
//	123
//	12
//	1
func Pattern6(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		row := ""
		for j := 1; j <= n-i; j++ {
			row += string(rune('0' + j%10))
		}
		b.add(row)
	}
	return b.String()
}

// Pattern7: star pyramid. Row i: (n-i-1) spaces, then (2i+1) stars.
//
//	  *
//	 ***
//	*****
func Pattern7(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		b.add(rep(" ", n-i-1) + rep("*", 2*i+1))
	}
	return b.String()
}

// Pattern8: inverted star pyramid. Row i: i spaces, then 2(n-i)-1 stars.
//
//	*****
//	 ***
//	  *
func Pattern8(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		b.add(rep(" ", i) + rep("*", 2*(n-i)-1))
	}
	return b.String()
}

// Pattern9: diamond = Pattern7 on top of Pattern8.
func Pattern9(n int) string {
	return Pattern7(n) + "\n" + Pattern8(n)
}

// Pattern10: half diamond. Stars grow 1..n, then shrink n-1..1.
//
//	*
//	**
//	***
//	**
//	*
func Pattern10(n int) string {
	var b builder
	for i := 1; i <= 2*n-1; i++ {
		stars := i
		if i > n {
			stars = 2*n - i
		}
		b.add(rep("*", stars))
	}
	return b.String()
}

// Pattern11: binary number triangle. Even rows start with 1, odd rows with 0,
// and the value flips at every step.
//
//	1
//	01
//	101
func Pattern11(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		v := 1
		if i%2 == 1 {
			v = 0
		}
		row := ""
		for j := 0; j <= i; j++ {
			row += string(rune('0' + v))
			v = 1 - v
		}
		b.add(row)
	}
	return b.String()
}

// Pattern12: number crown. Numbers 1..i, then 2*(n-i) spaces, then i..1.
//
//	1    1
//	12  21
//	123321
func Pattern12(n int) string {
	var b builder
	for i := 1; i <= n; i++ {
		left, right := "", ""
		for j := 1; j <= i; j++ {
			left += string(rune('0' + j%10))
		}
		for j := i; j >= 1; j-- {
			right += string(rune('0' + j%10))
		}
		b.add(left + rep(" ", 2*(n-i)) + right)
	}
	return b.String()
}

// Pattern13: increasing number triangle (Floyd's triangle).
// One counter keeps going across rows. Numbers are separated by spaces.
//
//	1
//	2 3
//	4 5 6
func Pattern13(n int) string {
	var b builder
	num := 1
	for i := 1; i <= n; i++ {
		parts := make([]string, 0, i)
		for j := 0; j < i; j++ {
			parts = append(parts, itoa(num))
			num++
		}
		b.add(strings.Join(parts, " "))
	}
	return b.String()
}

// Pattern14: increasing letter triangle.
//
//	A
//	AB
//	ABC
func Pattern14(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		row := ""
		for c := 'A'; c <= 'A'+rune(i); c++ {
			row += string(c)
		}
		b.add(row)
	}
	return b.String()
}

// Pattern15: reverse letter triangle.
//
//	ABC
//	AB
//	A
func Pattern15(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		row := ""
		for c := 'A'; c <= 'A'+rune(n-i-1); c++ {
			row += string(c)
		}
		b.add(row)
	}
	return b.String()
}

// Pattern16: alpha ramp. Row i repeats the i-th letter i+1 times.
//
//	A
//	BB
//	CCC
func Pattern16(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		b.add(rep(string(rune('A'+i)), i+1))
	}
	return b.String()
}

// Pattern17: alpha hill. Spaces, then letters going up to the middle and back down.
//
//	  A
//	 ABA
//	ABCBA
func Pattern17(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		row := rep(" ", n-i-1)
		breakpoint := (2*i + 1) / 2
		c := 'A'
		for j := 0; j < 2*i+1; j++ {
			row += string(c)
			if j < breakpoint {
				c++
			} else {
				c--
			}
		}
		b.add(row)
	}
	return b.String()
}

// Pattern18: alpha triangle. Row i starts at letter (n-1-i) and goes to the last letter.
//
//	C
//	BC
//	ABC
func Pattern18(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		row := ""
		for c := 'A' + rune(n-1-i); c <= 'A'+rune(n-1); c++ {
			row += string(c)
		}
		b.add(row)
	}
	return b.String()
}

// Pattern19: symmetric void. Top half: stars shrink with a growing gap.
// Bottom half: mirror image.
//
//	******
//	**  **
//	*    *
//	*    *
//	**  **
//	******
func Pattern19(n int) string {
	var b builder
	for i := 0; i < n; i++ { // top
		b.add(rep("*", n-i) + rep(" ", 2*i) + rep("*", n-i))
	}
	for i := 0; i < n; i++ { // bottom
		b.add(rep("*", i+1) + rep(" ", 2*(n-i-1)) + rep("*", i+1))
	}
	return b.String()
}

// Pattern20: symmetric butterfly.
//
//   - *
//     **  **
//     ******
//     **  **
//   - *
func Pattern20(n int) string {
	var b builder
	for i := 1; i <= 2*n-1; i++ {
		stars := i
		if i > n {
			stars = 2*n - i
		}
		b.add(rep("*", stars) + rep(" ", 2*(n-stars)) + rep("*", stars))
	}
	return b.String()
}

// Pattern21: hollow square. Stars only on the border.
//
//	****
//	*  *
//	*  *
//	****
func Pattern21(n int) string {
	var b builder
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				row += "*"
			} else {
				row += " "
			}
		}
		b.add(row)
	}
	return b.String()
}

// Pattern22: concentric number square of size (2n-1).
// Each cell = n - (distance to the nearest edge).
//
//	33333
//	32223
//	32123
//	32223
//	33333
func Pattern22(n int) string {
	var b builder
	size := 2*n - 1
	for i := 0; i < size; i++ {
		row := ""
		for j := 0; j < size; j++ {
			top, left, right, bottom := i, j, size-1-j, size-1-i
			d := min(min(top, bottom), min(left, right))
			row += itoa(n - d)
		}
		b.add(row)
	}
	return b.String()
}

// itoa converts a small non-negative int to a string without importing strconv.
func itoa(x int) string {
	if x == 0 {
		return "0"
	}
	s := ""
	for x > 0 {
		s = string(rune('0'+x%10)) + s
		x /= 10
	}
	return s
}
