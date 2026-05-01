// nums :   1 , 2 ,  4,   6
// prefix:  1 , 1 ,  2,   8
// postfix: 48, 24,  6,   1

// res.     48 , 24, 12 , 8



func productExceptSelf(nums []int) []int {

    prefix := make([]int, len(nums))
    postfix := make([]int, len(nums))
    res := make([]int, len(nums))

    prefix[0], postfix[len(nums)-1] = 1, 1
    for i:=1; i<len(nums); i++ {
        prefix[i] = nums[i-1]*prefix[i-1]
    }

    for i:= len(nums)-2; i>=0; i-- {
        postfix[i] = nums[i+1]*postfix[i+1]
    }

    for i:=0;i<len(res);i++ {
        res[i] = prefix[i] * postfix[i]
    }

    return res

}
