func canFinish(numCourses int, prerequisites [][]int) bool {
	ajlist := make([][]int, numCourses)
    for _, prerequisite := range prerequisites {
		ajlist[prerequisite[1]] = append(ajlist[prerequisite[1]],prerequisite[0])
	}

	visit := make(map[int]int)

	for i:=0; i<numCourses; i++ {
		if containLoop(i, ajlist, visit) {
			return false
		}
	}

	return true
	
    
}

func containLoop(course int, ajlist [][]int, visit map[int]int) bool {
	if visit[course] == 1 {
		return true
	}

	if len(ajlist[course]) == 0 {
		return false
	}

	visit[course] = 1

	for _, next := range ajlist[course] {
		if containLoop(next, ajlist, visit) {
			return true
		}
	}

	visit[course] = 0

	ajlist[course] = []int{}
	return false
   

}
