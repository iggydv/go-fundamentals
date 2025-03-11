package main

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r >= rows || r < 0 || c >= cols || c < 0 || grid[r][c] == '0' {
			return
		}
		// mark visited
		grid[r][c] = '0'

		dfs(r+1, c) // down
		dfs(r-1, c) // up
		dfs(r, c+1) // right
		dfs(r, c-1) // left
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count += 1
				dfs(r, c)
			}
		}
	}

	return count
}
