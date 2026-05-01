func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{})

    for _, num := range nums {
        if _, found := m[num]; found {
            return true
        }

        m[num] = struct{}{}
    }

    return false
    
}
