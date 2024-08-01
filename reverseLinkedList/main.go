package main

import (
	"fmt"
)

// ListNode определяет структуру узла односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList разворачивает односвязный список
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp

	}
	return prev
}

// printList печатает все значения узлов в списке
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// функция для создания примера списка и выполнения его разворота
func main() {
	// создание примера списка 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// печать оригинального списка
	fmt.Print("Original list: ")
	printList(head)

	// разворот списка
	reversedHead := reverseList(head)

	// печать развернутого списка
	fmt.Print("Reversed list: ")
	printList(reversedHead)
}
