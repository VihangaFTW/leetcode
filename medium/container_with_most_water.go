package medium

func ContainerWithMaxWater(height []int) int {

	maxArea := 0
	n := len(height)

	left, right := 0, n-1

	//? worst case time is O(N) cuz we examine each index exactly once
	for left != right {

		bottleneckHeight := min(height[left], height[right])
		//? area of rectange =  height *  width
		cur_area := bottleneckHeight * (right - left)
		// update max area as required
		maxArea = max(cur_area, maxArea)
		//! area = w*l. w keeps decrementing so we need to ensure that height evaluates to the biggest height to give us the max area
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
