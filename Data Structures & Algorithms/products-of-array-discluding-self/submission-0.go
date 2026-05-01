// nums :   1 , 2 ,  4,   6
// prefix:  1 , 2 ,  8,  48
// postfix: 48, 48,  24,  6

// res.     48 , 24, 12 , 8


// nums    :  -1,  0,  1,   2,  3
// prefix. :  -1,  0,  0 , 0,  0
// postfix :  0,   0,  6 ,  6 , 3
// res:       0,   -6  

func productExceptSelf(nums []int) []int {

    prefix := make([]int, len(nums))
    postfix := make([]int, len(nums))
    res := make([]int, len(nums))

    for i:=0; i<len(nums); i++ {
        if i == 0 {
            prefix[i] = nums[i]
            continue
        }

        prefix[i] =  nums[i]*prefix[i-1]

    }

    for i:= len(nums)-1; i>0; i-- {
        if i == len(nums)-1 {
            postfix[i] = nums[i]
            continue
        }

        postfix[i] = nums[i]*postfix[i+1]
    }

    for i:=0;i<len(res);i++ {
        if i == 0 {
            res[i] = 1 * postfix[i+1]
            continue
        }

        if i==len(res)-1 {
            res[i] = 1 * prefix[i-1]
            continue
        }

        res[i] = prefix[i-1] * postfix[i+1]
    }

    return res

}
