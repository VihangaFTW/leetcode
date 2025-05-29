package hard

//? Intuiton
// *for each column, the amount of water it can trap is determined by the minimum column size of the
//* largest column to the left and right of the current column.

//? Why this works?
//! max water a container can hold: two edges with the least height difference.
//! So we need to track the bigggest column to the left and right side for each column.

//*  the minimum optimal height of the two global max heights is basically the "bottleneck" and represents the maximum
//* water that can be held by the current column witthout overflowing. If we subtract the size of the current column size
//* from this minimum optimal, we basically remove the trapped water space taken by the column blocks and get how much
// *actual water is trapped by the current column.

//?  Time complexity : O(N) where N is the number of columns. Loop over entire list once. Space: O(1) due to inplace two pointer iteration

func TrappingRainWater(height []int) int {
	n := len(height)
	left, right := 0, n-1
	maxLeft, maxRight := height[left], height[right]

	trappedWater := 0

	for left < right {
		//! NOTE: the maxRight pointer might NOT point to the actual max right column as inner columns have not been traversed yet
		//! but, since we are taking the MINIMUM of the two maxes, we always know that maxLeft is ALWAYS smaller than/equals to maxRight as
		//! maxLeft has traversed through all the columns to the left side of current column. Because we start looping left -> right.
		//? smaller max determines the max amount of water that can be trapped. So we always shift the corresponding pointer to keep track of the latest bottleneck column
		if maxLeft < maxRight {
			//? at the start, left points to first column which can never trap water. So we ignore processing this column
			left++
			// left column smaller
			//? check if current column is new max. This is required to avoid "negative" trapped water in calculation below.
			maxLeft = max(maxLeft, height[left])
			trappedWater += maxLeft - height[left]

		} else {
			//? at the start, right points to last column which can never trap water. So we ignore processing this column
			right--
			// right column smaller
			maxRight = max(maxRight, height[right])
			trappedWater += maxRight - height[right]

		}
	}

	return trappedWater

}
