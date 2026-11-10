package mergeksortedlists

import (
	"container/heap"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

type Node = list.Node

// nodeHeap is a min-heap of list nodes ordered by value.
type nodeHeap []*Node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*Node)) }
func (h *nodeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// MergeKLists merges k sorted lists with a min-heap of heads. O(N log k).
func MergeKLists(lists []*Node) *Node {
	h := &nodeHeap{}
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}
	dummy := &Node{}
	tail := dummy
	for h.Len() > 0 {
		n := heap.Pop(h).(*Node)
		tail.Next = n
		tail = n
		if n.Next != nil {
			heap.Push(h, n.Next)
		}
	}
	return dummy.Next
}
