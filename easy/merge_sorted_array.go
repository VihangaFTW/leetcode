package easy

import "fmt"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	//* two pointers approach with an extra pointer to keep track of where to insert the next number
	//*  O(m+n) in worst case time as we traverse m+n slots max

	if m == 0 {
		copy(nums1, nums2)
		fmt.Println(nums1)
		return
	} else if n == 0 {
		fmt.Println(nums1)
		return
	}

	left, right, target := m-1, n-1, m+n-1

	for left >= 0 && right >= 0 {
		if nums1[left] > nums2[right] {
			nums1[target] = nums1[left]
			left--
		} else {
			nums1[target] = nums2[right]
			right--
		}

		target--
	}

	// Copy any remaining elements from nums2
	for right >= 0 {
		nums1[target] = nums2[right]
		right--
		target--
	}

	// Note: No need to copy remaining elements from nums1
	// because they are already in the correct position

	fmt.Println(nums1)

}
