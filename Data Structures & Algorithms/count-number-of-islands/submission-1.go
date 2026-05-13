var directions = [][]int{
	{0,1},
	{1,0},
	{0,-1},
	{-1,0},
}


func numIslands(grid [][]byte) int {
    numIslands := 0 
    
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j] == '1' {
                dfs(grid,i,j)
				numIslands++
			}
		}
	}

	return numIslands
}

func dfs(grid [][]byte,i,j int){
	for i<0 || j<0 || i>= len(grid) || j>=len(grid[0]) || grid[i][j] != '1' {
		return
	}

    grid[i][j] = '0'
    
	for _, direction := range directions {
        dfs(grid,i+direction[0],j+direction[1])
	}

}
