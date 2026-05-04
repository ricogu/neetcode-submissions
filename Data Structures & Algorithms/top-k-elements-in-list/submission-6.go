
func topKFrequent(nums []int, k int) []int {
    m:= make(map[int]int)

	for _, num := range nums {
		m[num] ++
	}

	freqBucket := make([][]int, len(nums)+1)

	for num, freq := range m {
		freqBucket[freq] = append(freqBucket[freq], num)
	}

    var res []int

	for i:=len(freqBucket)-1;i>0;i--{
        for _, num := range freqBucket[i] {
			res = append(res,num)
			if len(res) == k {
				return res
			}
		}
	}

	return res

}
