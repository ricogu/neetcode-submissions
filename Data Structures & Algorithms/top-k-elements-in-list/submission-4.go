import "slices"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v:= range nums {
		m[v] ++
	}

	var keys []int

	for key := range m {
		keys = append(keys,key)
	}

	slices.SortFunc(keys, func(a,b int) int{
		return m[b] - m[a]
	})

	return keys[:k]

}
