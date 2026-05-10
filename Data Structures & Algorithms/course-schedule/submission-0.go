func canFinish(numCourses int, prerequisites [][]int) bool {
	ajlist := make([][]int, numCourses)

	for _, prerequisite := range prerequisites{
		ajlist[prerequisite[1]] = append(ajlist[prerequisite[1]],prerequisite[0] )
	}

	visited := make(map[int]int)

    for i:=0; i<numCourses; i++ {
		if !dfs(i,ajlist, visited) {
			return false
		} 
	}

	return true
    
}

func dfs(course int, ajlist [][]int, visited map[int]int) bool {
    if visited[course] == 1 {
		return false
	}

	if len(ajlist[course]) == 0 {
		return true
	} 

	visited[course] = 1

	for _, aj := range ajlist[course] {
		if !dfs(aj,ajlist, visited) {
			return false
		}
	}

	visited[course] = 0

	return true

}
