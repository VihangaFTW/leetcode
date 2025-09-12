package medium

func HouseRobber2(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	// either we rob first house or last house since they are connected
	excludeFirstLoot := HouseRobber(nums[:len(nums)-1])
	excludeLastLoot := HouseRobber(nums[1:])
	return max(excludeFirstLoot, excludeLastLoot)

}
