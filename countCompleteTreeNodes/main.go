package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{Val: 1}
	head.Next = &TreeNode{Val: 2}
	head.Next.Next = &TreeNode{Val: 2}
	head.Next.Next.Next = &TreeNode{Val: 3}
	head.Next.Next.Next.Next = &TreeNode{Val: 4}
	head.Next.Next.Next.Next.Next = &TreeNode{Val: 2}
	result := removeElements(head, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	countNodes()

}
func countNodes(root *TreeNode) int {
	var count int
	for root != nil && root.Right != nil {
		count++
		root.Right = root.Right.Right
	}
	return count

}
