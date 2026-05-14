func findOrder(numCourses int, prerequisites [][]int) []int {
    ajlist := make([][]int, numCourses)
	inDegrees := make([]int,numCourses)
    for _, prerequisite := range prerequisites {
		ajlist[prerequisite[1]] = append(ajlist[prerequisite[1]],prerequisite[0])
		inDegrees[prerequisite[0]]++
	}

    var stack []int 
	var path  []int
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			stack = append(stack, course)
		}

	}

	for len(stack) > 0{
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		path = append(path, current)

        for _, next := range ajlist[current] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				stack = append(stack, next)
			}
		}
	}

	for _, inDegree := range inDegrees {
         if inDegree > 0 {
			return []int{}
		 } 
	}

	return path
}


