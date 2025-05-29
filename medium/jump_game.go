package medium

func JumpGame(nums []int) bool {

	n := len(nums)
	goal := n - 1

	for i := n - 2; i >= 0; i-- {
		//* jump index found that reaches current goal
		if i+nums[i] >= goal {
			// move goal closer
			goal = i
		}
		//* no closer goal found; might be in future iterations

	}
	//? if end can be reached, goal should point to the first index
	return goal == 0
}
