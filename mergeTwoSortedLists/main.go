package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     need:=&Listnode{}
//     for

// }

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{Val: 9}
	fmt.Println(l2)

}
