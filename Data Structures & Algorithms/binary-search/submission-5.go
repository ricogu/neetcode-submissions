func search(nums []int, target int) int {
	return binarySearch(nums,0,len(nums)-1, target)

}

func binarySearch(nums []int, start int, end int, target int) int{
	mid := (start+end)/2

	if start > end {
		return -1
	}

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return binarySearch(nums, mid+1, end, target)
	}

	return binarySearch(nums, start, mid-1, target)
}

