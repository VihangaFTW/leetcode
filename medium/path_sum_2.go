package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum2(root *TreeNode, targetSum int) [][]int {

	var paths [][]int

	var dfs func(node *TreeNode, curPath []int, remainder int) []int

	dfs = func(node *TreeNode, curPath []int, remainder int) []int {

		// base case: leaf node encountered
		if node.Left == nil && node.Right == nil {
			// path sum == target sum
			if remainder-node.Val == 0 {
				curPath = []int{node.Val}
				
			}
			// path sum != target sum
			return curPath
		}

		// inductive case: dfs both left and right
		lPath := dfs(node.Left, nil, remainder - node.Val)

		if lPath != nil {
			curPath = append(curPath, node.Val)
			return curPath
		}
		rPath := dfs(node.Right, nil, remainder - node.Val)

		if rPath != nil {
			curPath = append(curPath, node.Val)
			return curPath
		}

		return curPath
		
	}

	dfs(root, nil, targetSum)

	return paths
}
