package hard

func UniquePaths3(grid [][]int) int {

	totalPaths, totalECells, nRows, nCols := 0, 0, len(grid), len(grid[0])
	visited := make([][]int, len(grid))

	for row := range nRows {
		visited[row] = make([]int, nCols)
	}

	var dfs func(row int, col int, visited [][]int, eCellsFound int)
	dfs = func(row int, col int, visited [][]int, eCellsFound int) {

		// base case: hit an obstacle
		if grid[row][col] == -1 {
			return
		}

		// base case: reached exit with all empty cells visited
		if grid[row][col] == 2 && eCellsFound == totalECells {
			totalPaths += 1
			return
		}

		// mark cell as visited and increment count
		visited[row][col] = 1
		eCellsFound += 1

		// dfs up
		if row-1 >= 0 && visited[row-1][col] == 0 {
			dfs(row-1, col, visited, eCellsFound)
		}
		// dfs down
		if row+1 < nRows && visited[row+1][col] == 0 {
			dfs(row+1, col, visited, eCellsFound)
		}
		// dfs left
		if col-1 >= 0 && visited[row][col-1] == 0 {
			dfs(row, col-1, visited, eCellsFound)
		}
		// dfs right
		if col+1 < nCols && visited[row][col+1] == 0 {
			dfs(row, col+1, visited, eCellsFound)
		}

		// backtrack: unmark cell
		visited[row][col] = 0
		eCellsFound += 1

	}

	// find start position and count empty cells
	startRow, startCol := -1, -1

	for row := range nRows {
		for col := range nCols {
			if grid[row][col] == 1 {
				startRow, startCol = row, col
			}

			if grid[row][col] != -1 {
				totalECells += 1
			}

		}
	}

	// dfs from start
	dfs(startRow, startCol, visited, 1)

	return totalPaths

}
