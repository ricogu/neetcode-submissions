func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for idx, num := range nums{
		complement := target - num

		if idx_complement, found := m[complement]; found{
			return []int{idx_complement,idx}
		}

        m[num] = idx

	}

	return []int{0,0}
}
