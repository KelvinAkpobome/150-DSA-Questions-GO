package graphs

func NumberOfIslands() int8 {

	grid := [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "1", "0", "0"},
		{"0", "0", "0", "1", "1"},
	}

	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	rows := len(grid)
	cols := len(grid[0])

	var dfs func(grid [][]string, r int, c int)

	dfs = func(grid [][]string, r int, c int) {

		if r < 0 || c < 0 || c >= cols || r >= rows {
			return
		}

		if grid[r][c] == "2" || grid[r][c] == "0" {
			return
		}

		grid[r][c] = "2"

		for i := range 4 {
			nr := r + dirs[i][0]
			nc := c + dirs[i][1]
			dfs(grid, nr, nc)
		}
	}

	var count int8

	for r := range rows {
		for c := range cols {

			if grid[r][c] == "2" || grid[r][c] == "0" {
				continue
			}

			count++
			dfs(grid, r, c)

		}
	}

	return count
}
