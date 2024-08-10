package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	result := removeElements(head, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next

		} else {
			current = current.Next
		}
	}
	return head

}
