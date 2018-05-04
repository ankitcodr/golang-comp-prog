package leetcode

func explore(grid [][]int, r int, c int) int {
	if r < 0 || c < 0 || r > len(grid)-1 || c > len(grid[0])-1 || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	return 1 + explore(grid, r-1, c) + explore(grid, r, c-1) + explore(grid, r+1, c) + explore(grid, r, c+1)
}

func maxAreaOfIsland(grid [][]int) int {
	islandArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			exploredArea := explore(grid, i, j)
			if islandArea < exploredArea {
				islandArea = exploredArea
			}
		}
	}
	return islandArea
}

func MaxAreaOfIsland(grid [][]int) int {
	return maxAreaOfIsland(grid)
}
