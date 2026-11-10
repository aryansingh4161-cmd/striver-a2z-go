package minstack

import (
	"math/rand"
	"slices"
	"testing"
)

type minStacker interface {
	Push(int)
	Pop()
	Top() int
	GetMin() int
}

func TestMinStacks(t *testing.T) {
	for name, s := range map[string]minStacker{"pairs": &MinStack{}, "encoded": &MinStackEncoded{}} {
		var ref []int
		r := rand.New(rand.NewSource(8))
		for step := 0; step < 2000; step++ {
			if len(ref) == 0 || r.Intn(3) > 0 {
				x := r.Intn(41) - 20
				s.Push(x)
				ref = append(ref, x)
			} else {
				s.Pop()
				ref = ref[:len(ref)-1]
			}
			if len(ref) > 0 {
				if s.Top() != ref[len(ref)-1] || s.GetMin() != slices.Min(ref) {
					t.Fatalf("%s: step %d top=%d min=%d", name, step, s.Top(), s.GetMin())
				}
			}
		}
	}
}
