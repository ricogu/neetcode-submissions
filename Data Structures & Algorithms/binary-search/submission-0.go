func search(nums []int, target int) int {
    return binarySearch(0,len(nums)-1, nums,target)
}

func binarySearch(left, right int, nums[]int, target int) int {
    if left == right {
        if nums[left] == target {
            return left
        }
        return -1
    }

    mid := (left + right)/2

    leftFound := binarySearch(left,mid,nums,target)
    rightFound := binarySearch(mid+1,right,nums,target)

    if leftFound != -1 {
        return leftFound
    }

    return rightFound
    


}
