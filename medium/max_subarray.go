package medium

// MaxSubArray finds the maximum sum of any contiguous subarray within the given array.
// This implementation uses Kadane's algorithm which runs in O(n) time complexity.
//
// The algorithm works by maintaining two variables:
// - prevSum: the maximum sum ending at the current position
// - maxSum: the overall maximum sum found so far
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSubArray(nums []int) int {
	// Initialize both variables with the first element.
	// prevSum tracks the best sum ending at current position.
	// maxSum tracks the global maximum sum found so far.
	prevSum, maxSum := nums[0], nums[0]

	// Iterate through the array starting from the second element.
	for i := 1; i < len(nums); i++ {
		// Decide whether to extend the previous subarray or start a new one.
		// Check if appending nums[i] gives us a bigger optimal subarray sum.
		if potentialSum := prevSum + nums[i]; potentialSum > nums[i] {
			// Extending the previous subarray is beneficial.
			prevSum = potentialSum
		} else {
			// Starting a new subarray from current element is better.
			prevSum = nums[i]
		}

		// Update the global maximum if current subarray sum is better.
		if prevSum > maxSum {
			maxSum = prevSum
		}
	}

	return maxSum
}

// MaxSubArrayRecursive finds the maximum sum of any contiguous subarray using divide and conquer.
// This approach divides the array into halves and considers three cases:
// 1. Maximum subarray is entirely in the left half
// 2. Maximum subarray is entirely in the right half
// 3. Maximum subarray crosses the middle (spans both halves)
//
// Time Complexity: O(n log n)
// Space Complexity: O(log n) due to recursion stack
func MaxSubArrayRecursive(nums []int) int {
	return divideAndConquer(nums, 0, len(nums)-1)
}

// divideAndConquer recursively finds the maximum subarray sum in the given range [left, right].
func divideAndConquer(nums []int, left, right int) int {
	// Base case: single element.
	if left == right {
		return nums[left]
	}

	// Find the middle point to divide the array.
	mid := (left + right) / 2

	// Recursively find maximum subarray sum in left half.
	leftMax := divideAndConquer(nums, left, mid)

	// Recursively find maximum subarray sum in right half.
	rightMax := divideAndConquer(nums, mid+1, right)

	// Find maximum subarray sum that crosses the middle.
	crossMax := maxCrossingSubarray(nums, left, mid, right)

	// Return the maximum of all three cases.
	return max(leftMax, max(rightMax, crossMax))
}

// maxCrossingSubarray finds the maximum sum of subarray that crosses the middle point.
// It extends from middle to left and from middle to right, then combines both sums.
func maxCrossingSubarray(nums []int, left, mid, right int) int {
	// Find maximum sum extending from mid to left.
	leftSum := nums[mid]
	sum := nums[mid]
	for i := mid - 1; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// Find maximum sum extending from mid+1 to right.
	rightSum := nums[mid+1]
	sum = nums[mid+1]
	for i := mid + 2; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// Return the sum of both sides (crossing subarray).
	return leftSum + rightSum
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
