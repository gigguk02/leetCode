package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 3}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 0}
	d := &ListNode{Val: 4}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b

	fmt.Println(hasCycle(a))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
