// nums :   1 , 2 ,  4,   6
// prefix:  1 , 1 ,  2,   8
// postfix: 48, 24,  6,   1

// res:     1, 1, 2, 8

// res:          , 24 , 8



func productExceptSelf(nums []int) []int {

    res := make([]int, len(nums))

    for i, _ := range res{
        res[i] = 1
    }
   
    prefix := 1
    for i:=0; i<len(nums); i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i:= len(nums)-1; i>=0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
        
    }

    return res

}
