package medium

//! intuition
// * input list is always sorted. Hence, we can use two pointers starting at the
//* edge of the list and keep decrementing either one of them per iteration
//* until their sum hits the target.

func TwoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		curSum := numbers[left] + numbers[right]
		//? only one solutione exists and we found it
		if curSum == target {
			return []int{left + 1, right + 1}
			//? current sum is greater so we need to decrease right so the overall sum next time will decrease and might hit the target
		} else if curSum > target {
			right--
			//? current sum is smaller than target so increment left to get a bigger number next iteration
		} else {
			left++
		}
	}

	return nil
}
