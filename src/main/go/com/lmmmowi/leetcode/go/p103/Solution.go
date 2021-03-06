// 103.二叉树的锯齿形层次遍历[https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/]
package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Left = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}

	arr := zigzagLevelOrder(root)
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(arr[i][j], ",")
		}
		fmt.Println()
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	arr := make([][]int, 0)

	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}

	leftToRight := true
	for n := queue.Len(); n > 0; n = queue.Len() {
		sub := make([]int, 0)

		for i := 0; i < n; i++ {
			element := queue.Front()
			queue.Remove(element)

			node := element.Value.(*TreeNode)
			sub = append(sub, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		if !leftToRight {
			l := len(sub)
			for k := 0; k < l/2; k++ {
				temp := sub[k]
				sub[k] = sub[l-k-1]
				sub[l-k-1] = temp
			}
		}

		arr = append(arr, sub)
		leftToRight = !leftToRight
	}

	return arr
}
