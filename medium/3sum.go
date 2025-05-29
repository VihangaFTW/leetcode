package medium

import (
	"sort"
)

// ? time complexity: O(N^2) with O(logN) auxilary space for recursion stack when sorting
func ThreeSum(nums []int) [][]int {
	res := [][]int{}

	//? sort the input in ascending order so that we can run twoSum and skip duplicate numbers in O(1) time
	sort.Ints(nums) //? time O(N log N)

	//? if first number in sorted list is greater than O then we can never get a triplet that sums to 0
	if nums[0] > 0 {
		return nil
	}

	n := len(nums)

	//? time O(N^2): outer loop iterates through triplet[i] and inner loop iterates thru triplet[j] and triplet[k] using twoSum in O(N) time
	for start := range nums {

		if start > 0 && (nums[start] == nums[start-1]) {
			continue
		}

		left, right := start+1, n-1

		for left < right {
			sum := nums[start] + nums[left] + nums[right]
			//? unique triplet found: add to list
			if sum == 0 {
				res = append(res, []int{nums[start], nums[left], nums[right]})
				//? increment left pointer while ignoring repeated numbers
				left++
				for nums[left] == nums[left-1] && (left < right) {
					left++
				}
			//? sum bigger than 0: decremebt right
			}else if sum > 0 {
				//* decrement right ingoring all repeating numbers
				right--
			//? sum smaller than 0: increment left
			} else {
				//* incrmement left ingoring all repeating numbers
				left++
			}
		}

	}

	return res
}
