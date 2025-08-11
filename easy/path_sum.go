package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// If it's a leaf node, check if the sum matches
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// Recurse on left and right with reduced target
	return PathSum(root.Left, targetSum-root.Val) ||
		PathSum(root.Right, targetSum-root.Val)
}
