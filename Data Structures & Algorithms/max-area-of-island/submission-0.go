var directions = [][]int{
	{0,1},
	{1,0},
	{0,-1},
	{-1,0},
}

func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0 
    
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j] == 1 {
                connectedArea := countArea(grid,i,j)
				if connectedArea > maxArea {
					maxArea = connectedArea
				}
			}
		}
	}

	return maxArea

}

func countArea(grid [][]int,i,j int) int{
	for i<0 || j<0 || i>= len(grid) || j>=len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

    grid[i][j] = 0
    connectedArea := 0
	for _, direction := range directions {
        connectedArea += countArea(grid,i+direction[0],j+direction[1])
	}

	
	return  connectedArea + 1

}
