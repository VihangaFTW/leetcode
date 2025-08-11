package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {

	root := &easy.TreeNode{
		Val: 5,
		Left: &easy.TreeNode{
			Val: 4,
			Left: &easy.TreeNode{
				Val:   11,
				Left:  &easy.TreeNode{Val: 7},
				Right: &easy.TreeNode{Val: 2},
			},
		},
		Right: &easy.TreeNode{
			Val:  8,
			Left: &easy.TreeNode{Val: 13},
			Right: &easy.TreeNode{
				Val:   4,
				Right: &easy.TreeNode{Val: 1},
			},
		},
	}

	fmt.Println(easy.PathSum(root, 22))

}
