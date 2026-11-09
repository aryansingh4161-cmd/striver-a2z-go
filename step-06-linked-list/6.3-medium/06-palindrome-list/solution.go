package palindromelist

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// IsPalindromeStack compares with a stack of values. Space O(n).
func IsPalindromeStack(head *Node) bool {
	var st []int
	for cur := head; cur != nil; cur = cur.Next {
		st = append(st, cur.Val)
	}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val != st[len(st)-1] {
			return false
		}
		st = st[:len(st)-1]
	}
	return true
}

// IsPalindrome reverses the second half in place, compares, then restores it.
// Time O(n), space O(1).
func IsPalindrome(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := reverse(slow.Next)
	ok := true
	for a, b := head, second; b != nil; a, b = a.Next, b.Next {
		if a.Val != b.Val {
			ok = false
			break
		}
	}
	slow.Next = reverse(second) // restore the input
	return ok
}

func reverse(head *Node) *Node {
	var prev *Node
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
