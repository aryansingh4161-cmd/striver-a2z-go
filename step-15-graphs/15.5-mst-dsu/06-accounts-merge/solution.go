package accountsmerge

import (
	"sort"

	disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"
)

// Merge combines accounts that share an email. Each result is
// [name, emails sorted...]. O(N log N), N = total emails.
func Merge(accounts [][]string) [][]string {
	d := disjointset.New(len(accounts))
	owner := map[string]int{}
	for i, acc := range accounts {
		for _, email := range acc[1:] {
			if j, ok := owner[email]; ok {
				d.Union(i, j)
			} else {
				owner[email] = i
			}
		}
	}
	groups := map[int][]string{}
	for email, i := range owner {
		root := d.Find(i)
		groups[root] = append(groups[root], email)
	}
	out := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		out = append(out, append([]string{accounts[root][0]}, emails...))
	}
	return out
}
