package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := &ListNode{}
	head := sentinel
	carry, val := 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		val = carry + getVal(l1) + getVal(l2)

		carry = val / 10
		val = val % 10

		head.Next = &ListNode{val, nil}
		head = head.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return sentinel.Next
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}
func main() {
	// Создаем примеры связных списков
	// Пример 1: 342
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 9}

	// Пример 2: 465
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{Val: 9}

	// Вызываем функцию для сложения двух чисел
	result := addTwoNumbers(l1, l2)

	// Выводим результат
	fmt.Println("Результат сложения:")
	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Println("nil")
}
