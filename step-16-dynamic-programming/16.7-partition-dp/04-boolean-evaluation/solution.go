package booleanevaluation

const mod = 1_000_000_007

// combine returns (ways true, ways false) for "left op right".
func combine(op byte, lT, lF, rT, rF int) (int, int) {
	switch op {
	case '&':
		return lT * rT % mod, (lT*rF + lF*rT + lF*rF) % mod
	case '|':
		return (lT*rT + lT*rF + lF*rT) % mod, lF * rF % mod
	default: // '^'
		return (lT*rF + lF*rT) % mod, (lT*rT + lF*rF) % mod
	}
}

// Recursive returns (true ways, false ways) for s[i..j]. Exponential.
func Recursive(s string) int {
	var f func(i, j int) (int, int)
	f = func(i, j int) (int, int) {
		if i == j {
			if s[i] == 'T' {
				return 1, 0
			}
			return 0, 1
		}
		t, fl := 0, 0
		for k := i + 1; k < j; k += 2 {
			lT, lF := f(i, k-1)
			rT, rF := f(k+1, j)
			a, b := combine(s[k], lT, lF, rT, rF)
			t, fl = (t+a)%mod, (fl+b)%mod
		}
		return t, fl
	}
	t, _ := f(0, len(s)-1)
	return t
}

// Tab stores (true, false) counts per range. O(n³).
func Tab(s string) int {
	n := len(s)
	T, F := make([][]int, n), make([][]int, n)
	for i := range T {
		T[i], F[i] = make([]int, n), make([]int, n)
	}
	for i := n - 1; i >= 0; i -= 2 { // n is odd, so i always lands on a symbol
		if s[i] == 'T' {
			T[i][i] = 1
		} else {
			F[i][i] = 1
		}
		for j := i + 2; j < n; j += 2 {
			for k := i + 1; k < j; k += 2 {
				a, b := combine(s[k], T[i][k-1], F[i][k-1], T[k+1][j], F[k+1][j])
				T[i][j] = (T[i][j] + a) % mod
				F[i][j] = (F[i][j] + b) % mod
			}
		}
	}
	return T[0][n-1]
}
