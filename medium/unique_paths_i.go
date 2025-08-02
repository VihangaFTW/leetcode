package medium

func UniquePaths1(m int, n int) int {
	//* dp cache: cache[row][col] = number of unique paths to reach (row, col) from start.
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	//? base case: first row and column cells will each have ONLY 1 uniue path to get to each cell
	//? as the robot can only move right and down from the starting position.
	//! if start ==  end then the robot is already at the end and counts as one path
	for col := range n {
		cache[0][col] = 1
	}

	for row := range m {
		cache[row][0] = 1
	}

	//? overlapping subproblem: unique paths for given (row, col) = cache[row][col-1] + cache[row-1][col]
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			cache[row][col] = cache[row][col-1] + cache[row-1][col]
		}
	}

	return cache[m-1][n-1]

}

func UniquePaths1Optimized(m int, n int) int {
	//* condense previous implementation by reusing one row
	cache := make([]int, n)
	//* base case: fill up the first row with 1s (this is row 0 of prev implementation)
	for col := range n {
		cache[col] = 1
	}

	// fill up the rest of the slots by replacing the values in the array
	for range m-1 {
		for col := 1; col < n; col++ {
			cache[col] = cache[col] + cache[col-1]
		}
	}

	return cache[n-1]

}
