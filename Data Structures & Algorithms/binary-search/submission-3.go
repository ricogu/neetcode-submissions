func search(nums []int, target int) int {
    return binarySearch(0,len(nums)-1, nums,target)
}

func binarySearch(left, right int, nums[]int, target int) int {
    if left > right {
        return -1
    }

    mid := (left + right)/2

    if nums[mid] == target {
        return mid
    }

    if nums[mid] > target {
        return binarySearch(left,mid-1,nums,target)
    }

    return binarySearch(mid+1,right,nums,target)
    


}
