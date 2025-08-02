package medium

func UniquePaths2(obstacleGrid [][]int) int {

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	nRows, nCols := len(obstacleGrid), len(obstacleGrid[0])

	// Create 2D DP cache (auto-initialized to 0s)
	cache := make([][]int, nRows)
	for row := range nRows {
		cache[row] = make([]int, nCols)
	}

	// Initialize first row and column (stop at first obstacle)
	for col := 1; col < nCols && obstacleGrid[0][col] != 1; col++ {
		cache[0][col] = 1
	}

	for row := 1; row < nRows && obstacleGrid[row][0] != 1; row++ {
		cache[row][0] = 1
	}

	// Fill remaining cells: paths = left + top (skip obstacles)
	for row := 1; row < nRows; row++ {
		for col := 1; col < nCols; col++ {
			if obstacleGrid[row][col] != 1 {
				cache[row][col] = cache[row][col-1] + cache[row-1][col]
			}
		}
	}

	return cache[nRows-1][nCols-1]

}

func UniquePaths2Optimized(obstacleGrid [][]int) int {

	nRows, nCols := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 || obstacleGrid[nRows-1][nCols-1] == 1 {
		return 0
	}

	// Space-optimized: use 1D array like unique paths 1
	cache := make([]int, nCols)

	// Initialize first row (stop at first obstacle)
	for col := 0; col < nCols && obstacleGrid[0][col] != 1; col++ {
		cache[col] = 1
	}

	for row := 1; row < nRows; row++ {
		for col := range nCols {
			if obstacleGrid[row][col] == 1 {
				cache[col] = 0
			} else if col > 0 {
				cache[col] += cache[col-1]
			}
		}
	}

	return cache[nCols-1]

}
