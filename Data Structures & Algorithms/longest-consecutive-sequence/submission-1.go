import "slices"

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

     slices.Sort(nums)
	 maxLength := 1

     currentLength := 1
	 for i:=1; i< len(nums);i++ {

		if nums[i] == nums[i-1] {
			continue
		}
		if (nums[i] - nums[i-1]) == 1 {
			currentLength ++
		} else {
			currentLength = 1
		}

		if currentLength > maxLength {
			maxLength = currentLength
		}
		
	 }

	 return maxLength
}
