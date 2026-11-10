package maximumxorqueries

import (
	"sort"

	maximumxor "github.com/aryansingh4161-cmd/striver-a2z-go/step-17-tries/05-maximum-xor"
)

// Brute scans the array for each query. O(q·n).
func Brute(nums []int, queries [][2]int) []int {
	out := make([]int, len(queries))
	for qi, q := range queries {
		out[qi] = -1
		for _, v := range nums {
			if v <= q[1] {
				out[qi] = max(out[qi], q[0]^v)
			}
		}
	}
	return out
}

// MaximizeXor answers queries offline in increasing m, growing a bit trie.
func MaximizeXor(nums []int, queries [][2]int) []int {
	a := append([]int(nil), nums...)
	sort.Ints(a)
	order := make([]int, len(queries))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool { return queries[order[i]][1] < queries[order[j]][1] })

	t := maximumxor.NewBitTrie()
	out := make([]int, len(queries))
	p := 0
	for _, qi := range order {
		x, m := queries[qi][0], queries[qi][1]
		for p < len(a) && a[p] <= m {
			t.Insert(a[p])
			p++
		}
		if p == 0 {
			out[qi] = -1 // nothing small enough yet
		} else {
			out[qi] = t.MaxXor(x)
		}
	}
	return out
}
