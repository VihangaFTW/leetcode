package medium

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)

	closestSum := nums[0] + nums[1] + nums[2]

	if n == 3 {
		return closestSum
	}

	sort.Ints(nums)

	for start, startNum := range nums {

		if (start > 0) && startNum == nums[start-1] {
			continue
		}

		left, right := start+1, n-1

		for left < right {

			sum := startNum + nums[left] + nums[right]
			currentDiff, closestDiff := abs(sum, target), abs(closestSum, target)

			//? exact sum; cannot find a closer sum
			if currentDiff == 0 {
				return sum
			}
			//? new closest sum found
			if currentDiff <= closestDiff {
				closestSum = sum
			}
			//* move pointers so that next sum is closer to the target
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum

}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
