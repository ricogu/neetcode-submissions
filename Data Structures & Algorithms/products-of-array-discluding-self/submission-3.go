// [1,1,2,8]
// [48,24,6,1]


// [1, -1, 0 ,0 ,0]
// [ 0, 6 , 6 ,3 ,1]


func productExceptSelf(nums []int) []int {
	prefix := make([]int,len(nums))
	postfix := make([]int,len(nums))
	res := make([]int,len(nums))

    prefix[0] = 1
	postfix[len(nums)-1] = 1
	for idx:= 1; idx < len(nums); idx++ {
		prefix[idx] = prefix[idx-1]*nums[idx-1]

	}

	for idx:= len(nums)-2; idx >=0 ; idx--{
		postfix[idx] = postfix[idx+1]*nums[idx+1]
	}

	for idx:= 0; idx < len(nums); idx++ {
		res[idx] = prefix[idx]*postfix[idx]

	}

	return res
}
