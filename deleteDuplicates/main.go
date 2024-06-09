package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 1}
	result := deleteDuplicates(l1)
	printList(deleteDuplicates(result))

}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val}
	currentNew := newHead
	currentOld := head.Next

	for currentOld != nil {
		if currentOld.Val != currentNew.Val {
			currentNew.Next = &ListNode{Val: currentOld.Val}
			currentNew = currentNew.Next
		}
		currentOld = currentOld.Next
	}

	return newHead
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
