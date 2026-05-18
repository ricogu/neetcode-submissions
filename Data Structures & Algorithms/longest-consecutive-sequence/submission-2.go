func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

    m := make(map[int]int)
	for _, num := range nums {
		m[num] = 1
	}

	max := 1

	for _, num := range nums {
		curr := 1
		for i:= 1; m[num-i] == 1;i++ {
           curr++
		}

		if curr > max {
			max= curr
		}
	}

	return max
  
}
