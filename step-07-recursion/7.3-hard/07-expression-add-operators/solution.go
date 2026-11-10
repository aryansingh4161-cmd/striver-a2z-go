package expressionaddoperators

import "strconv"

// AddOperators returns all expressions over num's digits that evaluate to target.
func AddOperators(num string, target int) []string {
	res := []string{}
	var f func(i int, expr string, value, last int)
	f = func(i int, expr string, value, last int) {
		if i == len(num) {
			if value == target {
				res = append(res, expr)
			}
			return
		}
		for j := i; j < len(num); j++ {
			if j > i && num[i] == '0' {
				break // no leading zeros
			}
			part := num[i : j+1]
			x, _ := strconv.Atoi(part)
			if i == 0 {
				f(j+1, part, x, x) // first number: no operator before it
				continue
			}
			f(j+1, expr+"+"+part, value+x, x)
			f(j+1, expr+"-"+part, value-x, -x)
			f(j+1, expr+"*"+part, value-last+last*x, last*x)
		}
	}
	f(0, "", 0, 0)
	return res
}
