
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num:= range nums {
		m[num] ++
	}

	buckets := make([][]int, len(nums)+1)

	for num,freq := range m {
		buckets[freq] = append(buckets[freq],num)	
	}

	var res []int

	for i := len(nums); k>0; i-- {
        res = append(res,buckets[i]...)
        
		k-= len(buckets[i])
	}

    return res

}
