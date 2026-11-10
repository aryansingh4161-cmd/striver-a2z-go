package minimumcoins

// Denominations are Indian currency, largest first.
var Denominations = []int{2000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

// MinCoins returns the coins used to pay v greedily (largest first).
func MinCoins(v int) []int {
	res := []int{}
	for _, d := range Denominations {
		for v >= d {
			v -= d
			res = append(res, d)
		}
	}
	return res
}

// GreedyCount returns how many coins greedy uses for any coin system
// (coins sorted largest first). Shows where greedy can fail.
func GreedyCount(coins []int, v int) int {
	count := 0
	for _, c := range coins {
		count += v / c
		v %= c
	}
	if v != 0 {
		return -1
	}
	return count
}
