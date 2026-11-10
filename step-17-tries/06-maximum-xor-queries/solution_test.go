package maximumxorqueries

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMaximizeXor(t *testing.T) {
	if got := MaximizeXor([]int{0, 1, 2, 3, 4}, [][2]int{{3, 1}, {1, 3}, {5, 6}}); !slices.Equal(got, []int{3, 3, 7}) {
		t.Errorf("got %v", got)
	}
	if got := MaximizeXor([]int{5, 2, 4, 6, 6, 3}, [][2]int{{12, 4}, {8, 1}, {6, 3}}); !slices.Equal(got, []int{15, -1, 5}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 300; it++ {
		nums := make([]int, 1+r.Intn(10))
		for i := range nums {
			nums[i] = r.Intn(64)
		}
		qs := make([][2]int, 1+r.Intn(10))
		for i := range qs {
			qs[i] = [2]int{r.Intn(64), r.Intn(64)}
		}
		if got, want := MaximizeXor(nums, qs), Brute(nums, qs); !slices.Equal(got, want) {
			t.Fatalf("%v %v: %v vs %v", nums, qs, got, want)
		}
	}
}
