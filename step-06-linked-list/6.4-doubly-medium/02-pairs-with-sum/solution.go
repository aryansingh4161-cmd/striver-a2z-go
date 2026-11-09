package pairswithsum

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type DNode = list.DNode

// Pairs returns all pairs summing to target in a sorted DLL.
// Two pointers from both ends. Time O(n), space O(1) besides the answer.
func Pairs(head *DNode, target int) [][2]int {
	res := [][2]int{}
	if head == nil {
		return res
	}
	right := head
	for right.Next != nil {
		right = right.Next
	}
	left := head
	for left != right && right.Next != left {
		sum := left.Val + right.Val
		switch {
		case sum == target:
			res = append(res, [2]int{left.Val, right.Val})
			left = left.Next
			if left == right {
				break
			}
			right = right.Prev
		case sum < target:
			left = left.Next
		default:
			right = right.Prev
		}
	}
	return res
}
