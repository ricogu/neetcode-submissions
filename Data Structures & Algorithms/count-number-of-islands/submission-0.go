var DIRECTION = [][]int{
    {1,0},
    {0,1},
    {-1,0},
    {0,-1},
}

func numIslands(grid [][]byte) int {
    count:=0
    for i:=0;i<len(grid);i++ {
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j] == '1' {
                dfs(grid,i,j)
                count++
            }
        }
    }

    return count
}

func dfs(grid [][]byte, i,j int) {
    if i<0 || j<0 || i>= len(grid) || j>= len(grid[0]) || grid[i][j]=='0' {
        return
    }

    grid[i][j] = '0'

    for _, direction := range DIRECTION {
        next_i := i + direction[0]
        next_j := j + direction[1]
        dfs(grid,next_i,next_j)
    }
}
